package controllers

import (
	"konsultanku-app/database"
	"konsultanku-app/database/functions"
	function "konsultanku-app/database/functions"
	"konsultanku-app/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func CreateStudentProfile(c *gin.Context) {

	var getJson map[string]interface{}
	if err := c.ShouldBindJSON(&getJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentData, err := functions.CreateStudentAccount(getJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully create mse account", "data": studentData})
	return
}

func UpdateStudentProfile(c *gin.Context) {

	studentID, _ := c.Cookie("UID")
	jsonData, err := InputJson(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "data": nil})
		return
	}

	tag, err := functions.GetTagByName(jsonData["tag_name"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "data": nil})
		return
	}
	jsonData["tag_id"] = tag.ID
	delete(jsonData, "tag_name")
	dataUpdate, err := functions.UpdateStudentProfile(studentID, jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully update data", "data": dataUpdate})
	return
}

func Problem(c *gin.Context) {

	var problems []models.Problem

	category := c.Query("category")
	if category != "" {
		problemsWithCategory, err := function.GetProblemWithTags(category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		problems = problemsWithCategory
	} else {
		allProblem, err := function.GetAllProblem()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		problems = allProblem
	}
	result := []map[string]interface{}{}
	for i := range problems {
		mse, _ := function.GetMseByID(problems[i].MseID)
		tag, _ := function.GetTagByID(problems[i].TagID)
		jsonData := map[string]interface{}{
			"id_problem": problems[i].ID,
			"mse": map[string]interface{}{
				"mse_id":    mse.ID,
				"mse_name":  mse.MseName,
				"mse_since": mse.MseSince,
			},
			"problem":         problems[i].Problem,
			"like":            problems[i].Like,
			"comment_count":   problems[i].CommentCount,
			"problem_created": problems[i].ProblemCreated,
			"tag_name":        tag.TagName,
		}
		result = append(result, jsonData)
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func ProblemByID(c *gin.Context) {

	var problem models.Problem
	var tag models.Tags
	var mse models.MseProfile

	problemID := c.Param("id")
	result := database.DB.First(&problem, "id = ?", problemID)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Problem not found"})
		return
	}

	database.DB.First(&mse, "id = ?", problem.MseID)
	database.DB.First(&tag, "id = ?", problem.TagID)

	jsonData := map[string]interface{}{
		"id_problem":      problem.ID,
		"like":            problem.Like,
		"comment_count":   problem.CommentCount,
		"problem":         problem.Problem,
		"problem_created": problem.ProblemCreated,
		"tag_name":        tag.TagName,
		"mse": map[string]interface{}{
			"mse_name":  mse.MseName,
			"mse_since": mse.MseSince,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": jsonData})
	return
}

func BuildTeam(c *gin.Context) {

	var getJson map[string]interface{}
	if err := c.ShouldBindJSON(&getJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentID, _ := c.Cookie("UID")
	studentProfile, err := functions.CreateTeam(studentID, getJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": nil})
		return
	}
	c.SetCookie("TID", studentProfile["team_name"].(string), 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully create team", "data": gin.H{"team_name": getJson["team_name"].(string)}})
	return
}
