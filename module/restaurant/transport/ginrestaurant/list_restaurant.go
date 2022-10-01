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

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.ListRestaurantStore(store)
		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
