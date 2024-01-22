package telegram

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dj-yacine-flutter/y-z-a/utils"
	"github.com/zelenin/go-tdlib/client"
)

var (
	TDlibClient *client.Client
	CCChannel   = make(chan utils.CC)
)

func Start() error {
	var err error

	conf, err := utils.LoadConfig()
	if err != nil {
		return err
	}

	if conf.Telegram.AppID <= 0 {
		return errors.New("to use telegram put the required AppID in the config file")
	}

	if len(conf.Telegram.AppHash) < 32 {
		return errors.New("to use telegram put the required AppHash in the config file")
	}

	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	var (
		apiId   int32 = 11012090
		apiHash       = "5fa11a0398b42a30a6a8d124df5df129"
	)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        false,
		UseChatInfoDatabase:    false,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Y-Z-A",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		log.Fatalf("SetLogVerbosityLevel error: %s", err)
		return err
	}

	TDlibClient, err = client.NewClient(authorizer)
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
		return err
	}

	return nil
}

func Close() {
	if TDlibClient != nil {
		TDlibClient.Stop()
	}
}

func Stream() {
	listener := TDlibClient.GetListener()
	defer listener.Close()

	for update := range listener.Updates {
		if update.GetType() == client.TypeUpdateNewMessage {
			message := update.(*client.UpdateNewMessage).Message
			fmt.Println("Message Type:", message.GetType())

			switch content := message.Content.(type) {
			case *client.MessageText:
				fmt.Println("Message ID:", message.Id)
				fmt.Println("Chat ID:", message.ChatId)
				if message.ChatId == -750385682 {
					fmt.Println("Text:", content.Text.Text)
					cc := utils.CC{
						CCNUM: "5110200003199389",
						YEAR:  2025,
						MONTH: 12,
						NAME:  "test card name",
						CVV:   123,
					}
					CCChannel <- cc
				}
			}
		}
	}
}
