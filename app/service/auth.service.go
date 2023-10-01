package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Ahmad940/dropify/app/model"
	"github.com/Ahmad940/dropify/pkg/util"
	"github.com/Ahmad940/dropify/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var invalidCred = "Invalid email or password"

// Login
func Login(param model.Auth) (string, error) {
	var user model.User

	err := db.DB.Where("username LIKE ?", param.UserName).First(&user).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return "", errors.New(invalidCred)
		}
		return "", err
	}

	// comparing the passwordMatched whether it match with what exist in the database
	if passwordMatched := util.CheckPasswordHash(param.Password, user.Password); passwordMatched {
		// generating token for user base
		token, err := util.GenerateToken(user.ID)
		if err != nil {
			return "", err
		}
		return token, err
	} else {
		return "", errors.New(invalidCred)
	}
}

// CreateAccount
func CreateAccount(param model.Auth) error {
	var user model.User

	err := db.DB.Where("username = ?", param.UserName).First(&user).Error
	if SqlErrorIgnoreNotFound(err) != nil {
		fmt.Println("Yo")
		return err
	}

	// checking if user is registered or not
	if (user != model.User{}) {
		return errors.New("Username in use")
	}

	// hashing user password
	password, err := util.HashPassword(param.Password)
	if err != nil {
		return err
	}

	if strings.Trim(param.Role, "") == "" {
		param.Role = "user"
	}

	err = db.DB.Create(&model.User{
		ID:       gonanoid.Must(),
		UserName: param.UserName,
		Password: password,
	}).Error

	return nil
}
