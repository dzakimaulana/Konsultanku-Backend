package controllers

import (
	"konsultanku-app/database/functions"
	function "konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMseProfile(c *gin.Context) {

	var getJson map[string]interface{}
	if err := c.ShouldBindJSON(&getJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mseData, err := functions.CreateMseAccount(getJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully create mse account", "data": mseData})

}

func AllComments(c *gin.Context) {

	idProblemCookie, err := c.Request.Cookie("id_problem")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cookie not found"})
		return
	}

	comments, err := function.GetSuitableComment(idProblemCookie.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	dataResult := []map[string]interface{}{}
	for i := range comments {

		team, _ := function.GetTeamByID(comments[i].TeamID.String())
		jsonData := map[string]interface{}{
			"id_comment": comments[i].ID,
			"team": map[string]interface{}{
				"team_id":   team.ID,
				"team_name": team.TeamName,
			},
			"comment":         comments[i].Comment,
			"comment_created": comments[i].CommentCreated,
		}
		dataResult = append(dataResult, jsonData)
	}

	c.JSON(http.StatusOK, gin.H{"data": dataResult})
	return

}

func CreateProblem(c *gin.Context) {

	var getJson map[string]interface{}
	if err := c.ShouldBindJSON(&getJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}

	mseID, _ := c.Cookie("UID")
	problemData, err := functions.CreateProblem(mseID, getJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully upload problem", "data": problemData})
	return
}

func SendOffer(c *gin.Context) {

	idMseCookie, err := c.Cookie("id_mse")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cookie not found"})
		return
	}

	idTeamParam := c.Param("id")
	if idTeamParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID parameter is required"})
		return
	}
	_, err = function.GetTeamByID(idTeamParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Team not found!"})
	}
	idTeam, err := uuid.Parse(idTeamParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sendCollaboration, err := function.SendOffer(idMseCookie, idTeam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully send offer: " + sendCollaboration.ID.String()})
	return
}
