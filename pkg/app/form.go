package app

import (
	"strings"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

// 接口校验

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

// 实现Error接口的 Error 方法
func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v) // 参数绑定和入参校验
	// 翻译具体错误
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}