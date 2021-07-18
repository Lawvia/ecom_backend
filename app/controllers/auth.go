package controllers

import (
	"fmt"
	"strings"
	"time"

	"ecom_backend/app/dto"
	"ecom_backend/app/models"
	"ecom_backend/app/modules"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

func verifyAuth(reqToken string, ipClient string) (*dto.Token, string) {
	tk := &dto.Token{}
	var message string
	config := new(modules.Configuration).GetConfig()
	token, err := jwt.ParseWithClaims(reqToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("TOKEN_PASSWORD")), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		message = "Invalid Authorization"
		return nil, message
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		message = "Invalid Token"
		return nil, message
	}

	//success verified token, verified IP Whitelist
	fmt.Println("IP Client ", ipClient)
	var cron *dto.WhitelistData
	cron, err_w := new(models.ModelWhitelistData).CheckList(ipClient)
	fmt.Println(err_w)
	if cron == nil {
		//for development testing, allow localhost
		if ipClient == "::1" || ipClient == "127.0.0.1" {
			return tk, ""
		}

		message = "IP is not allowed! "+ipClient
		return nil, message
	}

	return tk, ""
}

func Login(c *gin.Context){
	var login dto.LoginRequest
	c.Bind(&login)

	if login.Username != "" && login.Password != "" {
		userModel := new(models.ModelUserData)

		//check username exist
		var cron *dto.UserData
		cron, err := userModel.CheckUsername(login.Username)
		if err != nil {
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": err,
			})
			return
		}

		//compare
		err = bcrypt.CompareHashAndPassword([]byte(cron.Password), []byte(login.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "Invalid credentials!",
			})
			return
		}

		//Create JWT token with HS256 algorithm
		config := new(modules.Configuration).GetConfig()
		tk := &dto.Token{
			UserId: cron.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(config.GetString("TOKEN_PASSWORD")))

		c.JSON(200, gin.H{
			"error": 0,
			"status": "SUCCESS",
			"message": "Logged in",
			"token": tokenString,
		})
	}else{
		if login.Password == "" {
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "Password can't be empty",
			})
		}else{
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "No username / JSON Error",
			})
		}
	}

}

func RegisterUser(c *gin.Context){
	var login dto.LoginRequest
	c.Bind(&login)

	if login.Username != "" && login.Password != "" {
		userModel := new(models.ModelUserData)

		//check username exist
		var cron *dto.UserData
		cron, err := userModel.CheckUsername(login.Username)
		if cron == nil {
			//attempt to create new user
			hash, _ := bcrypt.GenerateFromPassword([]byte(login.Password), 14)
			hashedPassword := string(hash)

			usr := new(dto.UserData)
			usr.Username = login.Username
			usr.Password = hashedPassword
			idUser, err_insert := userModel.Insert(usr)
			fmt.Println(idUser)

			if err_insert == nil {
				c.JSON(200, gin.H{
					"error": 0,
					"status": "SUCCESS",
					"message": "User succesfully registered",
				})
			}else{
				c.JSON(200, gin.H{
					"error": 1,
					"status": "Error",
					"message": err_insert,
				})
			}
			
		}else{
			fmt.Println(err)
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "Username already taken!",
			})
		}		
		
	}else{
		if login.Password == "" {
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "Password can't be empty",
			})
		}else{
			c.JSON(200, gin.H{
				"error": 1,
				"status": "Error",
				"message": "No username / JSON Error",
			})
		}
	}

}

func VerifyToken(c *gin.Context){
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

	c.JSON(200, gin.H{
		"error": 1,
		"status": "SUCCESS",
		"message": verify.UserId,
	})
}

func CheckAppStatus(c *gin.Context) {
	var status string = "OK"
	var cron []dto.ItemsData
	cron, err := new(models.ModelItemsData).Select()
	if err != nil {
		fmt.Println(err)
		status = "NOT OK"
	}
	fmt.Println(cron);
	c.JSON(200, gin.H{
		"server_date": time.Now(),
		"db_status":  status,
		// "zitems_check": cron,
	})
}