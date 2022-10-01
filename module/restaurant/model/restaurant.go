package restaurantmodel

import (
	"errors"
	"project/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Title           string         `json:"title" gorm:"column:title;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Title           string         `json:"title" gorm:"column:title;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil

}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()

}

type RestaurantUpdate struct {
	Name   *string        `json:"name" gorm:"column:name;"`
	Addr   *string        `json:"addr" gorm:"column:addr;"`
	Title  *string        `json:"title" gorm:"column:title;"`
	Status *int32         `json:"status" gorm:"default:1; column:status;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo"`
	Cover  *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
