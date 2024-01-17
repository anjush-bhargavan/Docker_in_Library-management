package models

import "gorm.io/gorm"

//FeedBack model holds the detail of user feedback
type FeedBack struct{
	ID			uint64			`json:"feedback_id" gorm:"primaryKey"`
	UserID		uint64			`json:"user_id" gorm:"not null" validate:"required"`
	Subject		string			`json:"subject" gorm:"not null" validate:"required"`
	Content		string			`json:"content"`
} 

//CreateFeedback will create a feedback for user
func (user *User) CreateFeedback(feedback *FeedBack, db *gorm.DB) error {
	result := db.Create(&feedback)
	if result.Error != nil {
		return result.Error
	}
	return nil
}