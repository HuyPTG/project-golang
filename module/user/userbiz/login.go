package userbiz

import (
	"context"
	"project/common"
	"project/component/tokenprovider"
	"project/module/user/usermodel"
)

type LoginStorage interface {
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	//appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hash          Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hash Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hash:          hash,
		expiry:        expiry,
	}
}

// 1.Find user, email
// 2.Hash pass from input and compare with pass in db
// 3.Provide: issue JWT token for client
// 3.1.Access token and refresh token
// 4.Return token(s)

func (biz *loginBusiness) Login(context context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(context, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}
	passHashed := biz.hash.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
