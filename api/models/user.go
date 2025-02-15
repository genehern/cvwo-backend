package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type User struct {
	ID       int   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func CreateUser( user *User) error {
	var existingUser User
	if err := DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {

		return fmt.Errorf("username '%s' already taken", user.Username)
	} 

	if err := user.HashPassword(); err != nil {
		return err
	}

	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	log.Print(user)
	return nil
}

func ValidateUser( user *User) error{
	var existingUser User
	err := DB.Where("username = ?", user.Username).First(&existingUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err // handle database error
	}
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Username '%s' not found", user.Username)
	}
	if err:= bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err!= nil{
		return fmt.Errorf("Wrong Password")
	}
	user.ID = existingUser.ID
	return nil
}

func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}