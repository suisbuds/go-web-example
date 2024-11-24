package logger

/*
1. 日志分级
2. 日志标准化
3. 日志格式化和输出
4.日志分级输出
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

const (
	SINGLE = "single"
	FRAMES = "frames"
)

type Level int8

type Fields map[string]interface{}

type Logger struct {
	base    *log.Logger
	ctx     context.Context
	fields  Fields   // 公共字段
	callers []string // 调用堆栈
}

var levels = [...]string{
	"debug",
	"info",
	"warn",
	"error",
	"fatal",
	"panic",
}

func (l Level) getString() string {
	if l < DEBUG || l > PANIC {
		return "unknown"
	}
	return levels[l]
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{base: l}
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

	// 获取调用堆栈信息
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		// 获取函数信息
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return nl
}

// 完整的调用栈
func (l *Logger) WithCallersFrames() *Logger {
	minDepth := 1
	maxDepth := 25
	callers := []string{}

	pcs := make([]uintptr, maxDepth)
	// 获取调用堆栈的程序计数器 pcs
	depth := runtime.Callers(minDepth, pcs)
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

func (l *Logger) format(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.getString()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	for k, v := range l.fields {
		// 避免覆盖
		if _, ok := data[k]; !ok {
			data[k] = v
		}
	}

	return data
}

func (l *Logger) output(level Level, mode string, message string) {
	var nl *Logger
	switch mode {
	case SINGLE:
		nl = l.WithCaller(3)
	case FRAMES:
		nl = l.WithCallersFrames()
	}
	// 生成格式化的日志，转换为 JSON 字符串
	body, _ := json.Marshal(nl.format(level, message))
	content := string(body)
	switch level {
	case FATAL:
		nl.base.Fatal(content)
	case PANIC:
		nl.base.Panic(content)
	default:
		nl.base.Print(content)
	}
}

func (l *Logger) Log(level Level, mode string, v ...interface{}) {
	l.output(level, mode, fmt.Sprint(v...))
}

func (l *Logger) Logf(level Level, mode string, format string, v ...interface{}) {
	l.output(level, mode, fmt.Sprintf(format, v...))
}
