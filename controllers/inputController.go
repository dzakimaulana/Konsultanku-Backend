package controllers

import "github.com/gin-gonic/gin"

func InputJson(c *gin.Context) (getJson map[string]interface{}, err error) {
	if err := c.ShouldBindJSON(&getJson); err != nil {
		return getJson, err
	}
	return getJson, nil
}
