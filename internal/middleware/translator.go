package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hans"
	//chinese ana english pkg
	uni "github.com/go-playground/universal-translator"

	enTran "github.com/go-playground/validator/v10/translations/en"
	//specific translator
	zhTran "github.com/go-playground/validator/v10/translations/zh"
)

func ValidateTranslator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := uni.New(en.New(), zh_Hans.New(), zh.New()) //adding language pck to translator
		local := ctx.GetHeader("local")
		//get then local via header
		trans, _ := u.GetTranslator(local)
		//getting  validator
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch local {
			case "zh":
				_ = zhTran.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = enTran.RegisterDefaultTranslations(v, trans)
			default:
				//default using chinese translator
				_ = zhTran.RegisterDefaultTranslations(v, trans)
			}
			ctx.Set("trans", trans)
		}
		ctx.Next()
	}
}