/*
Copyright © 2021 Matthew Turner <matthew.m.turner@outlook.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("analyze accepts exactly 1 file path argument")
			return
		}
		// wif simulate 70000000
		// wif simulate 70_000_000
		request := args[0]

		// User can pass two formats,
		// 	1. int without any comma separators or decimal
		//	2. int with underscore to make it easier to read
		// TODO: Need to validate input when using underscore.
		tokens := strings.Split(request, "_")
		cleanRequest := strings.Join(tokens, "")
		simulatedSize, err := strconv.Atoi(cleanRequest)

		if err != nil {
			fmt.Printf("Unable to convert simulated size to integer")
			return
		}

		printPerformanceResults(int64(simulatedSize))

	},
}

func init() {
	rootCmd.AddCommand(simulateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
