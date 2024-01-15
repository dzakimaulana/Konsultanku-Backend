package functions

import (
	"errors"
	"konsultanku-app/database"
	"konsultanku-app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTeamByID(teamID string) (models.Team, error) {
	var team models.Team
	if err := database.DB.First(&team, "id = ?", teamID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return team, errors.New("Team not found")
		}
		return team, err
	}
	return team, nil
}

func CreateTeam(studenID string, teamData map[string]interface{}) (student map[string]interface{}, err error) {

	teamID := uuid.New()
	team := models.Team{
		ID:              teamID,
		TeamName:        teamData["team_name"].(string),
		TeamCreated:     CurrentTime(),
		InCollaboration: false,
	}
	if err := database.DB.Create(&team).Error; err != nil {
		return student, err
	}
	dataJson := map[string]interface{}{
		"team_id":   teamID,
		"is_leader": true,
	}
	_, err = UpdateStudentProfile(studenID, dataJson)
	if err != nil {
		return teamData, err
	}
	return teamData, nil
}
