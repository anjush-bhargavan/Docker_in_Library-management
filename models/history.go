package models

import (
	"time"

	"gorm.io/gorm"
)

//History model the data that rented
type History struct {
	ID        		 uint64			`json:"history_id" gorm:"primaryKey"`
	UserID           uint64			`json:"user_id" gorm:"not null"`
	BookID           uint64			`json:"book_id" gorm:"not null"`
	Status	 		 string			`json:"status" gorm:"not null;default:'delivered'"`
	RentedOn         time.Time		`json:"rented_on"`
	ReturnedOn       time.Time		`json:"returned_on"`
	Remarks          string			`json:"remarks"`
}

//FetchHistoryByID will get the history of user by id
func (user *User) FetchHistoryByID(val uint64, db *gorm.DB) (*[]History, error) {
	var history []History
	if err := db.Find(&history).Where("id = ?", val).Error; err != nil {
		return nil,err
	}
	return &history,nil
}