package controllers

import (
	"errors"
	"net/http"

	"github.com/anjush-bhargavan/library-management/config"
	// "github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
)

var (
	getEmail      = GetEmail
	updateUser    = globalUser.UpdateUser
	getMembership = globalUser.FetchUserByID
	getHistory    = globalUser.FetchHistoryByID
	getUserID     = GetUserID
)

// UserProfile handles to get profile page of user
func UserProfile(c *gin.Context) {
	// data,_:=c.Get("email")
	email, result := getEmail(c)
	if !result {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "email not found from context",
		})
		return
	}
	user, err := getUser(email, config.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{"status": "Success",
		"message": "Profile fetched succesfully",
		"data":    user,
	})
}

// ProfileUpdate handles the updates of userprofile
func ProfileUpdate(c *gin.Context) {
	email, result := getEmail(c)
	if !result {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "Useremail not found",
		})
		return
	}
	// var user models.User
	user, err := getUser(email, config.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "User not found",
		})
		return
	}
	user.UserName = user.FirstName + " " + user.LastName

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "Failed",
			"message": "Binding error",
			"data":    err.Error(),
		})
		return
	}

	if err := updateUser(config.DB); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "Failed",
			"message": "Updating error",
		})
		return
	}
	c.JSON(200, gin.H{"status": "Success",
		"message": "Profile updated succesfully",
		"data":    user,
	})
}

// ChangePassword function helps to change password
func ChangePassword(c *gin.Context) {
	email, result := getEmail(c)
	if !result {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "Useremail not found",
		})
		return
	}

	user, err := getUser(email, config.DB)
	globalUser = *user
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed",
			"message": "User not found",
		})
		return
	}
	type password struct {
		Old  string `json:"old_password"`
		New  string `json:"new_password"`
		CNew string `json:"confirm_password"`
	}
	var newPassword password
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "Failed",
			"message": "Binding error",
			"data":    err.Error(),
		})
		return
	}
	if err := checkPassword(user, newPassword.Old); err != nil {
		c.JSON(http.StatusAccepted, gin.H{"status": "Failed",
			"message": "Password not correct",
			"data":    err.Error(),
		})
		return
	}
	// if newPassword.New == "" {
	// 	c.JSON(http.StatusNoContent, gin.H{"status": "Failed",
	// 		"message": "Password empty",
	// 		"data":    nil,
	// 	})
	// 	return
	// }
	if newPassword.New != newPassword.CNew {
		c.JSON(http.StatusConflict, gin.H{"status": "Failed",
			"message": "Password mismatch",
			"data":    nil,
		})
		return
	}

	if err := hashPassword(user, newPassword.New); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed",
			"message": "Failed to hash password",
			"data":    err.Error(),
		})
		return
	}
	if err := updateUser(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed",
			"message": "Failed to hash password",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success",
		"message": "Successfully changed password",
		"data":    nil,
	})

}

// ViewHistory handles to show the history of book taken by user
func ViewHistory(c *gin.Context) {
	userID, result := getUserID(c)
	if result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	history, err := getHistory(userID, config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in getting history",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"status": "Success",
		"message": "History fetched succesfully",
		"data":    history,
	})
}

// ViewMyPlan shows the plan of user
func ViewMyPlan(c *gin.Context) {
	// userIDContext, _ := c.Get("user_id")
	userID, result := getUserID(c)
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	membership, err := getMembership(userID, config.DB)

	if err == errors.New("NoRecord") {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Success",
			"message": "You haven't taken membership",
			"data":    err.Error(),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in getting membership",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"status": "Success",
		"message": "Your membership",
		"data":    membership,
	})
}

/////////////////////////////////////////////

//GetEmail will get the email saved in the context
func GetEmail(c *gin.Context) (string, bool) {
	data, flag := c.Get("email")
	if !flag {
		return "", false
	}

	email, ok := data.(string)
	if !ok {
		return "", false
	}

	return email, true
}

//GetUserID will return the userid saved in context
func GetUserID(c *gin.Context) (uint64, bool) {
	data, flag := c.Get("user_id")

	if !flag {
		return 0, false
	}

	userID, ok := data.(uint64)
	if !ok {
		return 0, false
	}

	return userID, true
}
