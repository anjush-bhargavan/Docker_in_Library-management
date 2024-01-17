package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

//Membership model holds the record of subcription
type Membership struct {
	ID			    		uint64          `json:"membership_id" gorm:"PrimaryKey"`
	UserID          		uint64			`json:"user_id" gorm:"not null;unique;foreignKey:UserID"`
	RazorpaySubscriptionID	string			`json:"subscription_id" gorm:"not null;unique"`
	Plan 					string			`json:"plan" gorm:"not null"`
	IsActive				bool			`json:"is_active" gorm:"not null;default:true"`
	StartedOn       		time.Time		`json:"started_on" gorm:"not null"`
	ExpiresAt       		time.Time		`json:"expires_at" gorm:"not null"`

}

//FetchUserByID will fetch the membership detail of user by id
func (member *User) FetchUserByID(val uint64, db *gorm.DB) (*Membership, error) {
	var membership Membership
	if err := db.Where("user_id = ?", val).First(&membership).Error; err == gorm.ErrRecordNotFound {
		return nil,errors.New("NoRecord")
	} else if err != nil {
		return nil, err
	}
	return &membership,nil
}