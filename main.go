/*
Copyright © 2024 Dj-Yacine
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dj-yacine-flutter/y-z-a/browser"
	"github.com/dj-yacine-flutter/y-z-a/cmd"
	"github.com/dj-yacine-flutter/y-z-a/telegram"
)

func main() {
	resp, err := http.Get("https://www.youtube.com/")
	if err != nil {
		log.Fatalln("Failed to connect the Internet.")
	}

	if resp.StatusCode != 200 {
		log.Fatalln("Not Authorized to use the app.")
	}

	fmt.Printf("\u001b[38;5;125m\u001b[48;5;0m%s\u001b[0m\n", fmt.Sprintln(`
                                                                
    ooooo  oooo           ooooooooooo                o          
      888  88             88    888                 888         
        888     ooooooooo     888     ooooooooo    8  88        
        888                 888    oo             8oooo88       
       o888o              o888oooo888           o88o  o888o     
                                                                `))
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
