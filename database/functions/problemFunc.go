package functions

import (
	"errors"
	"konsultanku-app/database"
	"konsultanku-app/models"

	"github.com/google/uuid"
)

func CreateProblem(mseID string, problemData map[string]interface{}) (problem models.Problem, err error) {

	tagID, err := GetTagByName(problemData["tag_name"].(string))
	if err != nil {
		tagID.ID = int32(0)
	}

	employees, ok := problemData["employees"].(float64)
	if !ok {
		return problem, errors.New("Invalid conversion")
	}

	problemID := uuid.New()
	problem = models.Problem{
		ID:             problemID,
		MseID:          mseID,
		TagID:          tagID.ID,
		Like:           int64(0),
		CommentCount:   int64(0),
		Problem:        problemData["problem"].(string),
		Description:    problemData["description"].(string),
		ProblemCreated: CurrentTime(),
		Income:         problemData["income"].(string),
		Employees:      int32(employees),
		LastSale:       problemData["late_sale"].(string),
		MediaSocial:    problemData["media_social"].(string),
		Goals:          problemData["goals"].(string),
		Address:        problemData["address"].(string),
	}
	if err := database.DB.Create(&problem).Error; err != nil {
		return problem, err
	}
	return problem, nil
}

func GetAllProblem() ([]models.Problem, error) {

	var problems []models.Problem
	if err := database.DB.Find(&problems).Error; err != nil {
		return nil, err
	}
	return problems, nil
}

func GetProblemWithTags(category string) ([]models.Problem, error) {

	problems, err := GetAllProblem()
	if err != nil {
		return nil, err
	}

	var tag models.Tags
	if err := database.DB.Find(&tag, "tag_name = ?", category).Error; err != nil {
		return nil, err
	}

	database.DB.Find(&problems, "tag_id = ?", tag.ID)
	return problems, nil
}

func GetProblemByID(problemID string) (models.Problem, error) {

	var problem models.Problem
	if err := database.DB.First(&problem, "id = ?", problemID).Error; err != nil {
		return problem, err
	}
	return problem, nil
}

func UpdateCommentCount(problemID string) error {
	problem, err := GetProblemByID(problemID)
	if err != nil {
		return err
	}
	problem.CommentCount += 1
	database.DB.Save(&problem)
	return nil
}

func AddLike(problemID string) error {

	problem, err := GetProblemByID(problemID)
	if err != nil {
		return err
	}
	problem.Like += 1
	database.DB.Save(&problem)
	return nil
}
