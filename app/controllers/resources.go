package controllers

import (
	"strings"

	"ecom_backend/app/models"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context){
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": "No Authorization",
		})
		return
	}

	splitted := strings.Split(reqToken, " ")
	if len(splitted) != 2 {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": "Invalid Authorization",
		})
		return
	}
	tokenPart := splitted[1]

	_, err := verifyAuth(tokenPart, c.ClientIP())
	if err != "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": err,
		})
		return
	}

	//attempt request
	cron, err_g := new(models.ModelItemsData).Select()
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err,
		})
	}else{
		c.JSON(200, gin.H{
			"error": 0,
			"status":  "Success",
			"data": cron,
		})
	}	
}