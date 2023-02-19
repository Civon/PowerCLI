/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "errors"
	// "strconv"

	// "github.com/manifoldco/promptui"
	"github.com/Civon/PowerCLI/internal/utils"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/fatih/color"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		a := utils.ReadCmdInput("Please input first string")
		b := utils.ReadCmdInput("Please input string to compare with")
		// fmt.Printf("You choose %q\n", a+b)

		//TODO (pre- & trailing spacing tolerant)

		bgColorRed := color.New(color.BgRed).SprintFunc()
		bgColorGreen := color.New(color.BgGreen).SprintFunc()
		if a == b {
			fmt.Println(bgColorGreen("The two strings are equal. "))
			return
		} else {
			fmt.Println(bgColorRed("The two strings are not equal. ") + "(Red marks the former different, Geen marks the latter)")
		}
		// compare char by char
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(a, b, true)
		colorRed := color.New(color.FgRed).SprintFunc()
		colorGreen := color.New(color.FgGreen).SprintFunc()

		for _, diff := range diffs {
			switch diff.Type {
			case diffmatchpatch.DiffDelete:
				fmt.Printf(colorRed(diff.Text))
			case diffmatchpatch.DiffInsert:
				fmt.Printf(colorGreen(diff.Text))
			case diffmatchpatch.DiffEqual:
				fmt.Printf(diff.Text)
			}

		}
		fmt.Printf("\n")
	},
}

// TODO not support powercli diff stringA stringB?
func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func main() {
// 	a := "c3Rya5nMQo=" // Red
// 	b := "c3Rya5nsMQo=" // Green

// }
