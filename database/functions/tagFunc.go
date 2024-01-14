package functions

import (
	"errors"
	"konsultanku-app/database"
	"konsultanku-app/models"

	"gorm.io/gorm"
)

func GetTagByID(idTag int32) (models.Tags, error) {
	var tag models.Tags
	if err := database.DB.First(&tag, "id = ?", idTag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tag, errors.New("Team not found")
		}
		return tag, err
	}
	return tag, nil
}

func GetTagByName(tagName string) (tag models.Tags, err error) {
	if err := database.DB.First(&tag, "tag_name = ?", tagName).Error; err != nil {
		return tag, err
	}
	return tag, nil
}
