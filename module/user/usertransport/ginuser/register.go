package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
	"project/component/hasher"
	"project/module/user/userbiz"
	"project/module/user/usermodel"
	"project/module/user/userstorage"
)

func Register(appContext appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FirstName))
	}
}
