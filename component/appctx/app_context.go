package appctx

import (
	"gorm.io/gorm"
	"project/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secret         string
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secret string) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secret: secret}

}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}
