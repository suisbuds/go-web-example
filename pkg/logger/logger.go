package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Level int8

type Fields map[string]interface{}

type Logger struct {
	baseLogger *log.Logger
	ctx        context.Context
	fields     Fields   // 公共字段
	callers    []string // 调用堆栈
}

var levels = [...]string{
	"debug",
	"info",
	"warn",
	"error",
	"fatal",
	"panic",
}

func (l Level) String() string {
	if l < DEBUG || l > PANIC {
		return "unknown"
	}
	return levels[l]
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{baseLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l

	// 避免浅拷贝问题
	if l.fields != nil {
		nl.fields = make(Fields, len(l.fields))
		for k, v := range l.fields {
			nl.fields[k] = v
		}
	}

	if l.callers != nil {
		nl.callers = make([]string, len(l.callers))
		copy(nl.callers, l.callers)
	}

	return &nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	nl := l.clone()
	if nl.fields == nil {
		nl.fields = make(Fields)
	}
	for k, v := range f {
		nl.fields[k] = v
	}
	return nl
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := l.clone()
	nl.ctx = ctx
	return nl
}

func (l *Logger) WithCaller(skip int) *Logger {
	nl := l.clone()

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return nl
}

// 完整的调用栈
func (l *Logger) WithCallersFrames() *Logger {
	minCallerDepth := 1
	maxCallerDepth := 25
	callers := []string{}

	pcs := make([]uintptr, maxCallerDepth)
	// 获取调用堆栈的程序计数器
	depth := runtime.Callers(minCallerDepth, pcs)
	// 程序计数器转换为调用堆栈帧
	frames := runtime.CallersFrames(pcs[:depth])

	for {
		frame, more := frames.Next()
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}

	nl := l.clone()
	nl.callers = callers
	return nl
}
