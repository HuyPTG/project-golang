package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"project/common"
	"project/component/uploadprovider"
	"project/module/upload/transport/uploadmodel"
	"strings"
	"time"
)

type CreateImageStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBiz( provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadBiz {
	return &uploadBiz{
		provider: provider,
		imgStore: imgStore,
	}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := filepath.Ext(fileName) // => "img.jpg" => ".jpg"
	fileName := fmt.Sprintf("%d%s",time.Now().Nanosecond(), fileExt) // => 098098098930.jpg ---> take unique
	img, err := biz.provider.SaveFileUpload(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}
	img.Width := w
	img.Height := h
	// img.CloudName = "s3" // should be set in provider
	img.Extension := fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ",err)
		return 0, 0, err
	}
	return img.Height, img.Height, nil
}