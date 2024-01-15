package functions

import (
	"konsultanku-app/database"
	"konsultanku-app/models"

	"github.com/google/uuid"
)

func GetSuitableComment(idProblem string) ([]models.TeamComment, error) {

	var comments []models.TeamComment
	if err := database.DB.Where("problem_id = ?", idProblem).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func AddComment(inputJson map[string]interface{}) (resultJson map[string]interface{}, err error) {

	commentID := uuid.New()
	teamID, err := uuid.Parse(inputJson["team_id"].(string))
	if err != nil {
		return resultJson, err
	}
	problemID, err := uuid.Parse(inputJson["problem_id"].(string))
	if err != nil {
		return resultJson, err
	}
	if err := UpdateCommentCount(problemID.String()); err != nil {
		return resultJson, err
	}
	teamComment := models.TeamComment{
		ID:             commentID,
		TeamID:         teamID,
		ProblemID:      problemID,
		Comment:        inputJson["comment"].(string),
		CommentCreated: CurrentTime(),
	}
	if err := database.DB.Create(&teamComment).Error; err != nil {
		return resultJson, err
	}
	resultJson = inputJson
	return resultJson, nil
}
