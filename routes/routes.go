package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(users_routes *gin.Engine) {
	users_routes.POST("/user/register")
	users_routes.POST("/user/login")
	users_routes.POST("/admin/add-product")
	users_routes.GET("/user/product")
	users_routes.GET("/user/search")
}
