package controllers

import (
	"fmt"

	"ecom_backend/app/dto"
	"ecom_backend/app/models"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context){
	ipClient := c.ClientIP()
	//for development testing, allow localhost
	if ipClient != "::1" && ipClient != "127.0.0.1" {
		var cron *dto.WhitelistData
		cron, err_w := new(models.ModelWhitelistData).CheckList(ipClient)
		fmt.Println(err_w)
		if cron == nil {
			message := "IP is not allowed! "+ipClient
			c.JSON(200, gin.H{
				"error": 1,
				"status":  "Error",
				"message": message,
			})
			return
		}
	}
	
	//attempt request
	data, err_g := new(models.ModelItemsData).Select()
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err_g,
		})
	}else{
		c.JSON(200, gin.H{
			"error": 0,
			"status":  "Success",
			"data": data,
		})
	}	
}