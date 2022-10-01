package restaurantbiz

import (
	"context"
	restaurantmodel "project/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type findRestaurantStore struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantStore {
	return &findRestaurantStore{store: store}

}

func (biz *findRestaurantStore) FindRestaurant(
	context context.Context,
	condition map[string]interface{},
	id string,
) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
