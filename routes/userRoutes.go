package routes

import (
	user "github.com/anjush-bhargavan/library-management/controllers/user"
	"github.com/gin-gonic/gin"
)

// function to handle user side routes
func userRoutes(r *gin.Engine) {
	user.RegisterUserRoutes(r)
	

}
