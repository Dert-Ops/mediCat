package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email           string `gorm:"type:varchar(100);unique;not null" json:"email"`
	PasswordHash    string `gorm:"type:varchar(255);not null" json:"password"`
	ProfilePicture  string `gorm:"type:varchar(255)" json:"profile_picture"`
	FullName        string `gorm:"type:varchar(100);not null" json:"fullname"`
	Age             int    `json:"age"`
	Bio             string `json:"bio"`
	IsVerified      bool   `gorm:"default:false" json:"is_verified"`
	GithubAccount   string `gorm:"type:varchar(255)" json:"github_account"`
	LinkedinAccount string `gorm:"type:varchar(255)" json:"linkedin_account"`
	GoogleAccount   string `gorm:"type:varchar(255)" json:"google_account"`
	Job             string `gorm:"type:varchar(100)" json:"job"`
	FavEmail        string `gorm:"type:varchar(100)" json:"fav_email"`
	Location        string `gorm:"type:varchar(100)" json:"location"`
}
