/*
Copyright © 2024 Dj-Yacine
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dj-yacine-flutter/y-z-a/account"
	"github.com/dj-yacine-flutter/y-z-a/browser"
	"github.com/dj-yacine-flutter/y-z-a/cmd"
	"github.com/dj-yacine-flutter/y-z-a/pb/uspb"
	"github.com/dj-yacine-flutter/y-z-a/telegram"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Printf("\u001b[38;5;125m\u001b[48;5;0m%s\u001b[0m\n", fmt.Sprintln(`
                                                                
    ooooo  oooo           ooooooooooo                o          
      888  88             88    888                 888         
        888     ooooooooo     888     ooooooooo    8  88        
        888                 888    oo             8oooo88       
       o888o              o888oooo888           o88o  o888o     
                                                                `))

	conn, err := grpc.Dial("0.0.0.0:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to check the server")
	}

	uscl := uspb.NewUserServiceClient(conn)
	if err = account.Sign(uscl); err != nil {
		log.Panic(err)
		log.Fatal("failed to manage user account")
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		browser.Close()
		telegram.Close()
		os.Exit(1)
	}()

	cmd.Execute()
}