package validation

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterMobile(translator ut.Translator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", ValidateMobile)
		_ = v.RegisterTranslation("mobile", translator, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} must have a value!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString("^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\\d{8}$", mobile)
	if !ok {
		return false
	}
	return true
}
