/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"
	"sync"

	"github.com/dj-yacine-flutter/y-z-a/browser"
	"github.com/dj-yacine-flutter/y-z-a/utils"
	"github.com/spf13/cobra"
)

// ta3hmCmd represents the ta3hm command
var ta3hmCmd = &cobra.Command{
	Use:   "ta3hm",
	Short: "use this command to test card in list of websites all in the same time",
	Long:  "use this command to test card in list of websites all in the same time",
	Run: func(cmd *cobra.Command, args []string) {
		headless, err := cmd.Flags().GetBool("headless")
		cobra.CheckErr(err)

		conf, err := utils.LoadConfig()
		cobra.CheckErr(err)

		links, err := utils.LoadLinks()
		cobra.CheckErr(err)

		ccs, err := utils.LoadCC()
		cobra.CheckErr(err)

		wg := sync.WaitGroup{}
		for _, link := range links {
			if !strings.Contains(link, "http") {
				continue
			}
			driver, err := browser.Chrome(headless)
			cobra.CheckErr(err)
			window, err := driver.CurrentWindowHandle()
			cobra.CheckErr(err)

			err = driver.ResizeWindow(window, 800, 600)
			cobra.CheckErr(err)
			wg.Add(1)
			go func(link string) {
				for _, cc := range ccs {
					err = utils.Fill(driver, false, conf, cc, link)
					if err != nil {
						log.Printf("error: %s \nLink : %s", err.Error(), link)
						continue
					}
				}
				wg.Done()
			}(link)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(ta3hmCmd)

	ta3hmCmd.Flags().BoolP("headless", "H", false, "use this flag to use chrome without GUI.")
}
