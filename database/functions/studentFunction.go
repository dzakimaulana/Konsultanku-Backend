package functions

import (
	"konsultanku-app/database"
	"konsultanku-app/models"
)

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

func UpdateStudentProfile(student models.StudentProfile) (models.StudentProfile, error) {

	result := database.DB.Model(&models.StudentProfile{}).Updates(student)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}
