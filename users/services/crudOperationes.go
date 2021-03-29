package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/mayusGomez/xalex/users"
	"github.com/mayusGomez/xalex/users/storage"
)

func GetUser(id string, s storage.Storage) (users.User, error) {

	user, err := s.Get(id)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}
	return user, nil

}

func validateCreateUser(user *users.User) (string, error) {
	if user.ID == "" {
		return "Code invalid", errors.New("User must have ID")
	}
	if len(user.Name) == 0 {
		return "Name invalid", errors.New("User without Name")
	}
	if len(user.Email) == 0 {
		return "Name invalid", errors.New("User without Email")
	}

	return "", nil
}

func CreateUser(user *users.User, s storage.Storage) error {

	errDesc, err := validateCreateUser(user)
	if err != nil {
		log.Println("Error validating data:", err)
		return errors.New(errDesc)
	}

	err = s.Create(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
