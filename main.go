package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"project/common"
	"project/component/appctx"
	"project/component/uploadprovider"
	"project/middleware"
	"project/module/restaurant/transport/ginrestaurant"
	"project/module/upload/transport/uploadtransport/ginupload"
	"project/module/user/usertransport/ginuser"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Addr            string        `json:"addr" gorm:"column:addr;"`
	Title           string        `json:"title" gorm:"column:title;"`
	Logo            *common.Image `json:"logo" gorm:"column:logo;"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Addr   *string `json:"addr" gorm:"column:addr;"`
	Title  *string `json:"title" gorm:"column:title;"`
	Status *int32  `json:"status" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
func main() {

	//test := Restaurant{
	//	Id:    1,
	//	Name:  "Provincial",
	//	Addr:  "43 Truong Dinh",
	//	Title: "Nope",
	//}
	//jsByte, err := json.Marshal(test) //{"id":1,"name":"Provincial","addr":"43 Truong Dinh","title":"Nope","status":0}
	//log.Println(string(jsByte), err)
	//json.Unmarshal([]byte("{\"id\":2,\"name\":\"Provincial\",\"addr\":\"45 Truong Dinh\",\"title\":\"Nope\",\"status\":0}"), &test)
	//log.Println(test.Id)

	//dsn := "host=localhost user=postgres password=123 dbname=go port=5454 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := os.Getenv("POSTGRES")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// xem log DB
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appContext := appctx.NewAppContext(db, s3Provider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("static", "./static")

	v1 := r.Group("/api/v1")

	// POST /upload
	v1.POST("/upload", ginupload.UploadImage(appContext))

	restaurant := v1.Group("/restaurant")
	// POST /restaurant
	restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))

	// DELETE /restaurant
	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// POST /register
	v1.POST("/register", ginuser.Register(appContext))

	// POST /authenticate
	v1.POST("/authenticate", ginuser.Login(appContext))

	// GET/ profile
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	// GET / restaurant
	//restaurant.GET("/:id", func(context *gin.Context) {
	//	var id = context.Param("id")
	//	if err != nil {
	//		context.JSON(http.StatusBadRequest, gin.H{
	//			"err": err.Error(),
	//		})
	//	}
	//
	//	var data Restaurant
	//	if err := context.ShouldBind(&data); err != nil {
	//		context.JSON(http.StatusBadRequest, gin.H{
	//			"err": err.Error(),
	//		})
	//		return
	//	}
	//	db.Where("id = ?", id).First(&data)
	//	context.JSON(http.StatusOK, gin.H{
	//		"data": data,
	//	})
	//})

	// GET LIST / restaurant
	restaurant.GET("", ginrestaurant.ListRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//Create
	//newRes := Restaurant{Name: "Coffee", Addr: "456 Dien Bien Phu ", Title: "Nha Hang So 4",}
	//
	//if err := db.Create(&newRes).Error; err != nil {
	//	log.Printf(err.Error())
	//} else {
	//	log.Printf("Create Success")
	//}

	//var readRes Restaurant
	//// Get
	//if err := db.Where("status = ?", 1).First(&readRes).Error; err != nil {
	//	log.Printf(err.Error())
	//} else {
	//	log.Printf("Get Success", readRes)
	//}
	//// Update
	//var writeRes Restaurant
	//writeRes.Name = "Nha Hang Dien Bien Phu"
	//if err := db.Where("id = ?", "250c22c0-73d2-4415-8912-de1a22cb1eec").Updates(&writeRes).Error; err != nil {
	//	log.Printf(err.Error())
	//} else {
	//	log.Printf("Update Success", writeRes)
	//}
	//
	////Delete
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", "6ef60028-6170-43cb-9678-84aeaa382ef0").Delete(nil).Error; err != nil {
	//	log.Printf(err.Error())
	//} else {
	//	log.Printf("Delete Success", writeRes)
	//}
	//
	//// Update For Data 0,null and " "
	//var dataNull int32 = 0
	//updateWriteRes := RestaurantUpdate{Status: &dataNull}
	//if err := db.Where("id = ?", "762338f6-f8f5-4248-9126-bd1bd838e298").Updates(&updateWriteRes).Error; err != nil {
	//	log.Printf(err.Error())
	//} else {
	//	log.Printf("Update For Data 0 , null and  Success", writeRes)
	//}
}
