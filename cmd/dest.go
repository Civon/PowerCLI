/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Civon/PowerCLI/internal/prefly"
	"github.com/spf13/cobra"
)

// destCmd represents the dest command
var destCmd = &cobra.Command{
	Use:   "dest",
	Short: "衝鋒小尖兵，飛往未知的道路",
	Long:  "犧牲小我完成大我，獲取一切情報的小尖兵，將為您踏上不歸路",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		pf, err := prefly.NewPreflyee(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		pf.Go()

		domain := getDomain(pf.Dest)
		fmt.Println("Domain: "+ domain)
		fmt.Println("Destination URL:\n"+ pf.Dest)
	},
}

//-
//? replace with map()
func getDomain(s string) (res string) {
	withoutQuery := strings.Split(s, "/")[2]
	withoutSubDomain := strings.Split(withoutQuery, ".")[2:]
	res = strings.Join(withoutSubDomain, ".")
	return
}


func init() {
	rootCmd.AddCommand(destCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
