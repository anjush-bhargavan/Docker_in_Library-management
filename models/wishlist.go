package models

import (
	"errors"

	"gorm.io/gorm"
)

//Wishlist model contains the id of books in cart
type Wishlist struct {
	ID		uint64	`json:"wishlist_id" gorm:"primaryKey"`
	UserID  uint64 	`json:"user_id" gorm:"not null"`  
	BookID 	uint64	`json:"book_id" gorm:"not null"`
}

// func GetWishListItem(bookId,userId uint64,db *gorm.DB) error {
	
// 	return nil
// }

//CreateWishlist function will a record in wishlist table
func (user *User) CreateWishlist(bookID,userID uint64,db *gorm.DB) error {
	var existingWishlist Wishlist
	if err := db.Where("book_id = ? AND user_id = ?",bookID,userID).First(&existingWishlist).Error; err == nil {
		return errors.New("record_exists")
	}else if err != gorm.ErrRecordNotFound {
		return err
	}
	var wishlist Wishlist
	wishlist.UserID = userID
	wishlist.BookID = bookID
	result := db.Create(&wishlist)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


//GetWishlist function will return the given users wishlist books
func (user *User) GetWishlist(userID uint64, db *gorm.DB) ([]Book,error){
	var bookIds []uint64
	if err := db.Model(&Wishlist{}).Where("user_id = ?",userID).Pluck("BookID",&bookIds).Error; err != nil {
		return nil,err
	}
	var books []Book
	if len(bookIds) == 0 {
		return nil, errors.New("no_items")
	}
	for _,id := range bookIds{
		var book Book
		if err := db.First(&book,id).Error; err != nil {
			return nil,err
		}
		books=append(books, book)
	}
	return books,nil
}


//RemoveItem function will remove an item from the users wishlist
func (user *User) RemoveItem(bookID,userID uint64,db *gorm.DB) error {
	var book Wishlist
	if err := db.Where("book_id = ? AND user_id = ?",bookID,userID).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}