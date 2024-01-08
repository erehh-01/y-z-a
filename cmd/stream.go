/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/dj-yacine-flutter/y-z-a/browser"
	"github.com/dj-yacine-flutter/y-z-a/telegram"
	"github.com/dj-yacine-flutter/y-z-a/utils"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "filter the cc from telgram msg.",
	Long:  "listener for incoming telegram messages and filter the cc from them.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		err = telegram.Start()
		cobra.CheckErr(err)

		if telegram.TDlibClient == nil {
			cobra.CheckErr(errors.New("failed to start telegram client"))
		}

		conf, err := utils.LoadConfig()
		cobra.CheckErr(err)

		links, err := utils.LoadLinks()
		cobra.CheckErr(err)

		go telegram.Stream()

		wg := sync.WaitGroup{}
		for i, link := range links {
			if i >= conf.Browser.MaxWindows {
				break
			}

			if !strings.Contains(link, "http") {
				continue
			}

			driver, err := browser.Chrome(false)
			cobra.CheckErr(err)
			window, err := driver.CurrentWindowHandle()
			cobra.CheckErr(err)

			err = driver.ResizeWindow(window, 800, 600)
			cobra.CheckErr(err)
			wg.Add(1)
			go func(link string) {
				wd, err := utils.Window(driver, conf, link)
				if err != nil {
					log.Printf("error: %s \nLink : %s", err.Error(), link)
				}

				tab := false
				for cc := range telegram.CCChannel {
					if tab {
						err = utils.Fill(driver, true, conf, cc, link)
						if err != nil {
							log.Printf("error: %s \n", err.Error())
						}
					} else {
						err = utils.Stream(wd, conf, cc)
						if err != nil {
							log.Printf("error: %s \n", err.Error())
						}
						tab = true
					}
				}

				wg.Done()
			}(link)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)

}
