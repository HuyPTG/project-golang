package middleware

import (
	"github.com/gin-gonic/gin"
	"project/common"
	"project/component/appctx"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppErr); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}
				appErr := common.ErrInternal(err.(error))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()
		context.Next()
	}
}
