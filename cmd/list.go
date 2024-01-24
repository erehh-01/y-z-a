/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/dj-yacine-flutter/y-z-a/browser"
	"github.com/dj-yacine-flutter/y-z-a/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this command to test card in list of websites one by one.",
	Long:  "Use this command to test card in list of websites one by one.",
	Run: func(cmd *cobra.Command, args []string) {
		headless, err := cmd.Flags().GetBool("headless")
		cobra.CheckErr(err)

		driver, err := browser.Chrome(headless)
		cobra.CheckErr(err)
		window, err := driver.CurrentWindowHandle()
		cobra.CheckErr(err)

		err = driver.ResizeWindow(window, 1920, 1080)
		cobra.CheckErr(err)

		checkouts, err := utils.LoadCheckouts()
		cobra.CheckErr(err)

		ccs, err := utils.LoadCC()
		cobra.CheckErr(err)

		conf, err := utils.LoadConfig()
		cobra.CheckErr(err)

		for _, link := range checkouts {
			if !strings.Contains(link, "http") {
				continue
			}
			for _, cc := range ccs {
				err = utils.Fill(driver, true, conf, cc, link)
				if err != nil {
					log.Printf("error: %s \nLink : %s", err.Error(), link)
					continue
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("headless", "H", false, "use this flag to use chrome without GUI.")
}
