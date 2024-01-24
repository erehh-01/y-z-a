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
	Short: "Filter the cc from telgram msg.",
	Long:  "Listener for incoming telegram messages and filter the cc from them.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		conf, err := utils.LoadConfig()
		cobra.CheckErr(err)

		err = telegram.Start(conf)
		cobra.CheckErr(err)

		if telegram.TDlibClient == nil {
			cobra.CheckErr(errors.New("failed to start telegram client"))
		}

		checkouts, err := utils.LoadCheckouts()
		cobra.CheckErr(err)

		go telegram.Stream(conf)

		wg := sync.WaitGroup{}
		for i, link := range checkouts {
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
			go func(link string, i int) {
				tab := false
				now := false
				rip := i

				wd, err := utils.Window(driver, conf, link)
				if err != nil {
					log.Printf("error: %s \n", err.Error())
				}

				for cc := range telegram.CCChannel {

					lk := link
					if rip < len(checkouts) {
						lk = checkouts[rip]
						rip++
					} else {
						rip = 0
					}

					if now {
						err = utils.Fill(driver, tab, conf, cc, lk)
						if err != nil {
							log.Printf("error: %s \n", err.Error())
						}
					} else {
						err = utils.Stream(wd, conf, cc)
						if err != nil {
							log.Printf("error: %s \n", err.Error())
						}
						now = true
					}
					tab = true
				}

				wg.Done()
			}(link, i)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
}
