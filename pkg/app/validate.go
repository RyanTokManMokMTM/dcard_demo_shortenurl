package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

type ValidateErr struct {
	Key     string
	Message string
}

type ValidateErrs []*ValidateErr

//Error error interface implementation
func (v *ValidateErr) Error() string {
	return v.Message
}

//Error error interface implementation
func (v ValidateErrs) Error() string {
	return strings.Join(v.Errors(), ",")
}

//Errors separate all error message in the list
func (v ValidateErrs) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

//BindingAndValidating binding and validate
//Binding data from ctx to v
func BindingAndValidating(ctx *gin.Context, v interface{}) (bool, ValidateErrs) {
	//it may have lots of error with binding/validating
	var errs ValidateErrs
	err := ctx.ShouldBind(v)
	if err != nil {
		//translation and get translator from header
		v := ctx.Value("trans")

		//casting to translator
		trans, _ := v.(ut.Translator)
		errV, ok := err.(validator.ValidationErrors) //to validate error
		if !ok {
			return false, errs //
		}

		//translate the error with trans
		for key, value := range errV.Translate(trans) { //translate error from trans
			errs = append(errs, &ValidateErr{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil //not any error ,passed
}

func TimeValidation(fl validator.FieldLevel) bool {
	now := time.Now().Unix()

	//parse field time to time UTC ,then convert to unix time
	utcTime := "2006-01-02T03:04:05Z"
	expired, _ := time.Parse(utcTime, fl.Field().String())

	return expired.Unix() < now
}
