package uploadmodel

import (
	"errors"
	"project/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileToLarge = common.NewCustomError(
		errors.New("file to large"),
		"file to large",
		"ErrFileToLarge",
	)
)

func ErrCannotSaveFile(err error) *common.AppErr {
	return common.NewCustomError(
		err,
		"cannot save file",
		"ErrCannotSaveFile",
	)
}

func ErrFileIsNotImage(err error) *common.AppErr {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}
