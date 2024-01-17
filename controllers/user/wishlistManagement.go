package controllers

import (
	"net/http"
	"strconv"

	"github.com/anjush-bhargavan/library-management/config"
	// "github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	// "gorm.io/gorm"
)

var (
	createWishlist = globalUser.CreateWishlist
	getBook 	   = globalUser.GetBook 
	getWishlist    = globalUser.GetWishlist
	removeItem     = globalUser.RemoveItem
)
//AddToWishlist function add books to Wishlist
func AddToWishlist(c *gin.Context) {
	userID, result := getUserID(c)
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	stringID :=c.Param("id")
	id, err := strconv.ParseUint(stringID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest,gin.H{	"status":"Failed",
											"message":"Error parsing string",
											"data":err.Error(),
										})
        return
    }

	if _,err := getBook(id,config.DB); err != nil{
		c.JSON(http.StatusNotFound,gin.H{	"status":"Failed",
											"message":"Book not found in records",
											"data":err.Error(),
										})
		return
	}

	
	
	if err := createWishlist(id,userID,config.DB); err == errors.New("record_exist") {
		c.JSON(http.StatusConflict,gin.H{	"status":"Failed",
											"message":"Book already added to wishlist",
										})
		return
	}else if err !=  nil {
		c.JSON(http.StatusInternalServerError,gin.H{	"status":"Failed",
														"message":"Database error",
														"data":err.Error(),
													})
		return
	}
	
	
	c.JSON(http.StatusOK,gin.H{	"status":"Success",
								"message":"Book added to wishlist",
							})

}


//ShowWishlist function lists the Wishlist items
func ShowWishlist(c *gin.Context) {
	userID, result := getUserID(c)
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	books,err := getWishlist(userID,config.DB)
	if  err != nil {
		c.JSON(http.StatusBadGateway,gin.H{	"status":"Failed",
											"message":"Database error",
											"data":err.Error(),
										})
		return
	}

	c.JSON(http.StatusOK,gin.H{	"status":"Success",
								"message":"Books in wishlist :",
								"data":books,
							})
}


//DeleteWishlist function deletes the Wishlist items
func DeleteWishlist(c *gin.Context) {
	userID, result := getUserID(c)
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed",
			"message": "Error in userId",
		})
		return
	}

	stringID :=c.Param("id")
	bookID, err := strconv.ParseUint(stringID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest,gin.H{	"status":"Failed",
											"message":"Error parsing string",
											"data":err.Error(),
										})
        return
    }
	

	if err := removeItem(bookID,userID,config.DB); err != nil {
		c.JSON(http.StatusBadGateway,gin.H{	"status":"Failed",
											"message":"Error deleting from database",
											"data":err.Error(),
										})
		return
	}
	c.JSON(http.StatusOK,gin.H{	"status":"Success",
								"message":"Book removed from wishlist",
								"data":nil,
							})

}