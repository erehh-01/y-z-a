package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dj-yacine-flutter/y-z-a/pb/uspb"
	"github.com/dj-yacine-flutter/y-z-a/utils"
)

func (account *Account) createAccount(file, device string) error {
	var email string

	fmt.Println("Enter your email address:")
	for {
		if _, err := fmt.Scanln(&email); err != nil {
			fmt.Println("Error reading email:", err)
			continue
		}

		if !utils.IsValidEmail(email) {
			fmt.Println("Enter a valid email address:")
		} else {
			break
		}
	}

	var name string

	fmt.Println("Enter a username (must contain only lowercase letters, digits, or underscore and more then 5 characters):")
	for {
		if _, err := fmt.Scanln(&name); err != nil {
			fmt.Println("Error reading username:", err)
			continue
		}

		if !utils.IsValidName(name) {
			fmt.Println("Enter a valid username:")
		} else {
			break
		}
	}

	var password string

	fmt.Println("Enter a password (must contain more then 6 characters):")
	for {
		if _, err := fmt.Scanln(&password); err != nil {
			fmt.Println("Error reading password:", err)
			continue
		}

		if !utils.IsValidPassword(password) {
			fmt.Println("Enter a valid password:")
		} else {
			break
		}
	}

	res, err := account.uscl.CreateUser(context.Background(), &uspb.CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
		Device:   device,
	})
	if err != nil {
		log.Fatal("failed to create user account")
	}

	if res.User != nil {
		data := UserData{
			Name:     res.GetUser().GetName(),
			Email:    res.GetUser().GetEmail(),
			Password: password,
		}

		jsonData, err := json.Marshal(&data)
		if err != nil {
			return errors.New("failed to save data to the account.json")
		}

		os.WriteFile(file, jsonData, 0644)
		fmt.Printf("Hi, %s\n\n", res.GetUser().GetName())
	}
	return nil
}
