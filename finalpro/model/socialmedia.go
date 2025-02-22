package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Name           string `gorm:"not null;type:varchar(100)" json:"name" form:"name" valid:"required~Your social media name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media url is required,url~Please input valid url"`
	UserID         uint
	User           *User
}

type RequestSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

type ResponseSocialMedia struct {
	ID             uint   `json:"id" example:"1"`
	Name           string `json:"name" example:"Instagram"`
	SocialMediaURL string `json:"social_media_url" example:"https://www.sosmed.com/username"`
	User           struct {
		Username string `json:"username" example:"anon"`
		Email    string `json:"email" example:"test@example.com"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	DeletedAt gorm.DeletedAt
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
