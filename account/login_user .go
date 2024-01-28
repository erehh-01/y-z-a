package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/dj-yacine-flutter/y-z-a/pb/uspb"
	"github.com/dj-yacine-flutter/y-z-a/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (account *Account) loginAccount(file, device string) error {
	data := UserData{}
	dataBytes, err := os.ReadFile(file)
	if err != nil {
		return errors.New("error in account file")
	}

	var email string
	var password string

	if err = json.Unmarshal(dataBytes, &data); err != nil {
		fmt.Println("account file is empty. You will enter your data manualy !")
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

		fmt.Println("Enter your password:")
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
	} else {
		email = data.Email
		password = data.Password
	}

	res, err := account.uscl.LoginUser(context.Background(), &uspb.LoginUserRequest{
		Email:    email,
		Password: password,
		Device:   device,
	})
	if err != nil {
		switch status.Code(err) {
		case codes.PermissionDenied:
			return ErrPermissionDenied
		case codes.NotFound:
			return ErrUserNotFound
		default:
			return ErrUnkown
		}
	}

	if res.User != nil {
		data = UserData{
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
