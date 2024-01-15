package functions

import (
	"errors"
	"konsultanku-app/database"
	"konsultanku-app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func StudentRole(studentID string) bool {

	var student models.StudentProfile
	if err := database.DB.First(&student, "id = ?", studentID).Error; err != nil {
		return false
	}
	role := student.Role
	if role != "Mahasiswa" {
		return false
	}
	return true
}

func CreateStudentAccount(StudentData map[string]interface{}) (student models.StudentProfile, err error) {

	birthDate, err := DateConvert(StudentData["date_of_birth"].(string))
	if err != nil {
		return student, err
	}
	student = models.StudentProfile{
		ID:          StudentData["id_student"].(string),
		Role:        "Mahasiswa",
		StudentName: StudentData["student_name"].(string),
		DateOfBirth: birthDate,
		IsLeader:    false,
		Major:       StudentData["major"].(string),
		University:  StudentData["university"].(string),
		ClassOf:     StudentData["class_of"].(string),
	}
	if err := database.DB.Create(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func UpdateStudentProfile(studentID string, studentJson map[string]interface{}) (map[string]interface{}, error) {

	result := database.DB.Model(&models.StudentProfile{}).Where("id = ?", studentID).Updates(&studentJson)
	if result.Error != nil {
		return studentJson, result.Error
	}
	return studentJson, nil
}

func InTeam(studentID string) (teamID string, err error) {

	var student models.StudentProfile
	if err := database.DB.Where("id = ?", studentID).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return teamID, errors.New("User not found")
		}
		return teamID, err
	}

	inTeam := student.TeamID != uuid.Nil
	if !inTeam {
		return teamID, err
	}
	teamID = student.TeamID.String()
	return teamID, nil
}

func GetStudentByID() {

}
