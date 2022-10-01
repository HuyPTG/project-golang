package restaurantstorage

import (
	"context"
	"project/common"
	restaurantmodel "project/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	page *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&page.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (page.Page - 1) * page.Limit

	if err := s.db.
		Offset(offset).
		Limit(page.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
