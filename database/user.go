package database

import (
	"errors"

	"github.com/HanmaDevin/workoutdev/types"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *types.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return DB.Create(user).Error
}

func LoginUser(email, password string) (types.User, error) {
	var user types.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, errors.New("invalid password")
	}
	return user, nil
}
