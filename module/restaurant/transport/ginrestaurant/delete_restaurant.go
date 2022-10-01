package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
	restaurantbiz "project/module/restaurant/biz"
	restaurantstorage "project/module/restaurant/storage"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var id = c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing required code request part",
			})
			return
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
