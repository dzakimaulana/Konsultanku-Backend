package functions

import (
	"errors"
	"konsultanku-app/database"
	"konsultanku-app/models"

	"gorm.io/gorm"
)

func MseOnly(mseID string) bool {

	var mse models.MseProfile
	if err := database.DB.First(&mse, "id = ?", mseID).Error; err != nil {
		return false
	}
	role := mse.Role
	if role != "UMKM" {
		return false
	}
	return true
}

func CreateMseAccount(mseData map[string]interface{}) (mse models.MseProfile, err error) {

	mse = models.MseProfile{
		ID:        mseData["id_mse"].(string),
		Role:      "UMKM",
		OwnerName: mseData["owner_name"].(string),
		MseName:   mseData["mse_name"].(string),
		MseType:   mseData["mse_type"].(string),
		MseSince:  mseData["mse_since"].(string),
	}
	if err := database.DB.Create(&mse).Error; err != nil {
		return mse, err
	}
	return mse, nil
}

func UpdateMseAccount(mseData map[string]interface{}) (models.MseProfile, error) {
	mse := models.MseProfile{
		OwnerName: mseData["owner_name"].(string),
		MseName:   mseData["mse_name"].(string),
		MseType:   mseData["mse_type"].(string),
		MseSince:  mseData["mse_since"].(string),
	}
	if err := database.DB.Create(&mse).Error; err != nil {
		return mse, err
	}
	return mse, nil
}

func GetMseByID(idMse string) (models.MseProfile, error) {
	var mse models.MseProfile
	if err := database.DB.First(&mse, "id = ?", idMse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mse, errors.New("Team not found")
		}
		return mse, err
	}
	return mse, nil
}
