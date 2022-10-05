package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u := context.MustGet(common.CurrentUser)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
