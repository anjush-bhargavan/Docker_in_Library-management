package controllers

import (
	"net/http"

	"github.com/anjush-bhargavan/library-management/config"
	"github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
)

var createFeedback = globalUser.CreateFeedback

//Feedback handles users to give their feedback
func Feedback(c *gin.Context) {
	userID, result := getUserID(c)
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	var feedback models.FeedBack

	if err :=c.ShouldBindJSON(&feedback);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"Failed",
											"message":"Error while Binding",
											"data":err.Error(),
										})
		return
	}
	feedback.UserID=userID
	if err := validate.Struct(feedback); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{	"status":"Failed",
											"message":"Please fill all fields",
										})
		return
	}

	if err := createFeedback(&feedback,config.DB); err!= nil {
		c.JSON(http.StatusBadGateway,gin.H{"status":"Failed",
											"message":"Database error",
											"data":err.Error(),
										})
		return
	}
	c.JSON(200,gin.H{	"status":"Success",
						"message":"Feedback sent succesfully",
						"data":feedback,
					})

}