package gin

import (
	"strings"

	"github.com/coderi421/gframework/app/pkg/code"
	"github.com/coderi421/gframework/pkg/common/core"
	"github.com/coderi421/gframework/pkg/errors"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// 用来去除验证错误消息中的顶层结构体名称的 为了好看
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 接收 Gin 的上下文对象、错误对象和翻译器对象，用于处理验证错误
func HandleValidatorError(c *gin.Context, err error, trans ut.Translator) {
	errs, ok := err.(validator.ValidationErrors)

	if !ok {
		core.WriteResponse(c, errors.WithCode(code.ErrCodeInvalidParam, "Invalid Param: %s"), err.Error())
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
	}
	core.WriteResponse(c, errors.WithCode(code.ErrCodeInvalidParam, "Invalid Param: %s"), removeTopStruct(errs.Translate(trans)))
	//c.JSON(http.StatusBadRequest, gin.H{
	//	"error": removeTopStruct(errs.Translate(trans)),
	//})
	return
}
