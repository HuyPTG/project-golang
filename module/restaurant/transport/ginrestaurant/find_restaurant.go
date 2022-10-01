package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"project/common"
	restaurantbiz "project/module/restaurant/biz"
	restaurantstorage "project/module/restaurant/storage"
)

func FindRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing required code request part",
			})
			return
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.FindRestaurantStore(store)
		data, err := biz.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
