package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAddToWishlist(t *testing.T) {
	test := []tests{
		{
			name:        "success",
			route:       "/user/wishlist/2",
			errorResult: nil,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1, true
			}
			getBook = func(bookId uint64, db *gorm.DB) (*models.Book, error) {
				var book models.Book
				return &book, nil
			}
			createWishlist = func(bookId, userId uint64, db *gorm.DB) error {
				return nil
			}

			w,err := Setup(http.MethodPost,tc.route,nil,authToken)
			if err != nil {
				t.Fatal(err)
			}

			if tc.errorResult != nil {
				errValue , _ := json.Marshal(tc.body)
				require.JSONEq(t,w.Body.String(),string(errValue))
			}else{
				require.Equal(t,w.Code,200)
			}
		})
	}
}


func TestShowWishlist(t *testing.T) {
	test := []tests {
		{
			name: "success",
			route: "/user/wishlist",
			errorResult: nil,
		},
	}

	for _,tc := range test {
		t.Run(tc.name,func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1, true
			}
			getWishlist = func(userId uint64, db *gorm.DB) ([]models.Book, error) {
				var books []models.Book
				return books,nil
			}

			w,err := Setup(http.MethodGet,tc.route,nil,authToken)
			if err != nil {
				t.Fatal(err)
			}

			if tc.errorResult != nil {
				errValue,_ := json.Marshal(tc.errorResult)
				require.JSONEq(t,w.Body.String(),string(errValue))
			}else{
		
				require.Equal(t,w.Code,200)
			}
		})
	}
}


func TestDeleteWish(t *testing.T) {
	test := []tests {
		{
			name: "success",
			route: "/user/wishlist/2",
			errorResult: nil,
		},
	}
	
	for _,tc := range test {
		t.Run(tc.name , func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1, true
			}
			removeItem = func(bookId, userId uint64, db *gorm.DB) error {
				return nil
			}

			w,err := Setup(http.MethodDelete,tc.route,nil,authToken)
			if err != nil {
				t.Fatal(err)
			}

			if tc.errorResult != nil {
				errValue,_ := json.Marshal(tc.errorResult)
				require.JSONEq(t,w.Body.String(),string(errValue))
			}else {
				require.Equal(t,w.Code,200)
			}
		})
	}
}