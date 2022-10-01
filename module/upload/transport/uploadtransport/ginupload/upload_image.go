package ginupload

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/common"
	"project/component/appctx"
	"project/module/upload/transport/uploadbiz"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Upload file in local service
		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		//	Id:        0,
		//	Url:       "http://localhost:8080/static/" + fileHeader.Filename,
		//	Width:     0,
		//	Height:    0,
		//	CloudName: "local",
		//	Extension: "png",
		//}))
		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close() // We can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
