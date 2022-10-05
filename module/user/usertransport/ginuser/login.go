package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
	"project/component/hasher"
	"project/component/tokenprovider/jwt"
	"project/module/user/userbiz"
	"project/module/user/usermodel"
	"project/module/user/userstorage"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := context.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(context.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
