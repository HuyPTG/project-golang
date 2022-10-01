package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
	restaurantbiz "project/module/restaurant/biz"
	restaurantmodel "project/module/restaurant/model"
	restaurantstorage "project/module/restaurant/storage"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		//go func() {
		//	defer common.AppRecover()
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
