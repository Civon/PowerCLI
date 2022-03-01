/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go/types"
	"log"
	"regexp"

	"github.com/spf13/cobra"
)

// flattenUrlCmd represents the flattenUrl command
var flattenUrlCmd = &cobra.Command{
	Use:   "flatten",
	Short: "flatten a shorterned url",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var shorternedUrl = "https://t.ly/nkeY"

        if len(args) >= 1 && args[0] != "" {
            shorternedUrl = args[0]
		}
		// remove http/https if match
		re, err := regexp.Compile(`[http][s]?[://]`)
		if err != nil {
			log.Fatal(err)
			//? or
			fmt.Println(err)
		}
		shorternedUrl = re.ReplaceAllString(shorternedUrl, "")
		fmt.Println("Try to get '" + shorternedUrl + "' Info...")
		
		// data provied by WOT & unshorten.link
		URL := "https://unshorten.link/check?url=https://" + shorternedUrl
		
		// get dataset
		const (
			Undefined Safety = iota
			// Danger
			// Unknown
			Safe
		)

		type UrlMeta struct {
			redirectUrl	string
			safety int
			cookies int
			// ? Domain? if trust level is recognize as safety, still need domain check?   
			// trustLevel -> numbers from WOT in future
		}
		// the site is "Dangerous"!  [warning color]
		// with [how many] cookies
		// [links] [saperate dots if dangerous]
	},
}

func init() {
	rootCmd.AddCommand(flattenUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flattenUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flattenUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
