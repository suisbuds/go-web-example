package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations"github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

// 注册验证器注册，支持错误提示多语言
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uniTrans := ut.New(en.New(), zh.New(), ja.New())
		locale := c.GetHeader("locale")
		trans, _ := uniTrans.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			case "ja":
				_ = ja_translations.RegisterDefaultTranslations(v,trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}

		c.Next()
	}
}