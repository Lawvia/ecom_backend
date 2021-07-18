package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"crypto/rand"

	"ecom_backend/app/dto"
	"ecom_backend/app/models"
	// "ecom_backend/app/modules"
	"github.com/gin-gonic/gin"
)

func GetCartData(c *gin.Context){
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

	verify, err := verifyAuth(tokenPart, c.ClientIP())
	if err != "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": err,
		})
		return
	}

	//attempt request
	cron, err_g := new(models.ModelCartData).Select(verify.UserId)
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err,
		})
	}else{
		var grandTotal int
		for i:=0; i<len(cron); i++ {
			converted,_ := strconv.Atoi(cron[i].Price)
			grandTotal += converted * cron[i].Qty
		}
		c.JSON(200, gin.H{
			"error": 0,
			"status":  "Success",
			"data": cron,
			"grand_total": grandTotal,
		})
	}	
}

func UpdateCart(c *gin.Context){
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

	verify, err := verifyAuth(tokenPart, c.ClientIP())
	if err != "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": err,
		})
		return
	}

	//attempt request
	cartModel := new(models.ModelCartData)
	var cart dto.UpdateCartRequest
	c.Bind(&cart)
	fmt.Println("cart ", cart)

	cron, _ := new(models.ModelUserData).CheckUsername(verify.UserId)
	_, err_g := cartModel.PurgeCart(cron.UserId)
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err,
		})
		return
	}
	
	for i:=0; i<len(cart.Cart); i++ {
		newCart := new(dto.CartTable)
		newCart.IdItems = cart.Cart[i].IdItems
		newCart.IdUser = cron.UserId
		for j:=0; j<cart.Cart[i].Qty; j++ {
			_, err_i := cartModel.Insert(newCart)
			fmt.Println(err_i)
		}
	}
	c.JSON(200, gin.H{
		"error": 0,
		"status":  "Success",
	})
}

func GetTransaction(c *gin.Context){
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

	verify, err := verifyAuth(tokenPart, c.ClientIP())
	if err != "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": err,
		})
		return
	}

	//attempt request
	headerModel := new(models.ModelHeaderData)

	user, _ := new(models.ModelUserData).CheckUsername(verify.UserId)
	cron, err_g := headerModel.SelectHeader(user.UserId)
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err,
		})
	}else{
		var finalResp []dto.GetTransactionResp
		if len(cron) > 0 {
			for i :=0; i<len(cron); i++ {
				//get detail transaction
				var transactResp dto.GetTransactionResp
				transactResp.PurchaseKey = cron[i].PurchaseKey
				transactResp.Address = cron[i].Address
				transactResp.Total = cron[i].Total
				transactResp.Status = cron[i].Status
				
				det, err_d := headerModel.SelectDetail(cron[i].PurchaseKey)
				fmt.Println(err_d)
				transactResp.Detail = det

				finalResp = append(finalResp, transactResp)
			}
		}
		
		c.JSON(200, gin.H{
			"error": 0,
			"status":  "Success",
			"data": finalResp,
		})
	}	
}

func Checkout(c *gin.Context){
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

	verify, err := verifyAuth(tokenPart, c.ClientIP())
	if err != "" {
		c.JSON(403, gin.H{
			"error": 1,
			"status": "Error",
			"message": err,
		})
		return
	}

	//attempt request
	headerModel := new(models.ModelHeaderData)
	itemModel := new(models.ModelItemsData)
	var checkO dto.CheckoutRequest
	c.Bind(&checkO)

	user, _ := new(models.ModelUserData).CheckUsername(verify.UserId)
	cron, err_g := new(models.ModelCartData).Select(verify.UserId)
	if err_g != nil {
		c.JSON(200, gin.H{
			"error": 1,
			"status":  "Error",
			"message": err,
		})
	}else{
		//generate purchase key
		key := make([]byte, 8)
		_, err := rand.Read(key)
		if err != nil {
			//log error to files
		}
		purchaseKey := fmt.Sprintf("%x",key)

		var grandTotal int
		for i:=0; i<len(cron); i++ {
			converted,_ := strconv.Atoi(cron[i].Price)
			grandTotal += converted * cron[i].Qty

			//insert to detail
			newDetail := new(dto.DetailData)
			newDetail.IdItems = cron[i].IdItems
			newDetail.Qty = cron[i].Qty
			newDetail.HeaderId = purchaseKey

			_, err_i := headerModel.InsertDetail(newDetail)
			fmt.Println(err_i)

			//deplete stock from ms_items
			_, err_p := itemModel.UpdateStock(cron[i].IdItems, cron[i].Qty)
			fmt.Println(err_p)
		}
		//insert to header
		newHeader := new(dto.HeaderData)
		newHeader.PurchaseKey = purchaseKey
		newHeader.IdUser = user.UserId
		newHeader.Address = checkO.Address
		newHeader.Total = strconv.Itoa(grandTotal)

		_, err_i := headerModel.InsertHeader(newHeader)
		fmt.Println(err_i)

		//purge cart
		_, err_g := new(models.ModelCartData).PurgeCart(user.UserId)
		fmt.Println(err_g)

		c.JSON(200, gin.H{
			"error": 0,
			"status":  "Success",
		})
	}	
}