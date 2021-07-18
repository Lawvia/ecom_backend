package mappings

import (
	"ecom_backend/app/controllers"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/static"
)

var Router *gin.Engine

func CreateUrlMappings(){
	Router = gin.Default()
	// v1 of the API
	v1 := Router.Group("/api")
	{
		v1.GET("/app_status/", controllers.CheckAppStatus)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/verify_token", controllers.VerifyToken)
			auth.POST("/register", controllers.RegisterUser)
		}

		resource := v1.Group("/resource")
		{
			resource.GET("/get_items", controllers.GetItems)
		}

		transaction := v1.Group("/transaction")
		{
			transaction.GET("/get_cart", controllers.GetCartData)
			transaction.GET("/get_transaction", controllers.GetTransaction)
			transaction.POST("/update_cart", controllers.UpdateCart)
			transaction.POST("/checkout", controllers.Checkout)
		}
	}
}