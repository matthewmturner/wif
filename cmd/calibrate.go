/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// calibrateCmd represents the calibrate command
var calibrateCmd = &cobra.Command{
	Use:   "calibrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		wifSetup()
		runCalibration()

	},
}

func init() {
	rootCmd.AddCommand(calibrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calibrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calibrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func wifDirExists() bool {
	home, home_err := os.UserHomeDir()
	if home_err != nil {
		return false
	}
	wif_path := fmt.Sprintf("%s/.wif", home)
	_, wif_err := os.Stat(wif_path)
	if wif_err == nil {
		return true
	}
	if os.IsNotExist(wif_err) {
		return false
	}
	return false
}

func createWifDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	wif_path := fmt.Sprintf("%s/.wif", home)
	return os.MkdirAll(wif_path, os.ModePerm)
}

func wifSetup() {
	if wifDirExists() {

	} else {
		fmt.Println("~/.wif dir does not exist")
		err := createWifDir()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("wif dir created")
	}
}

func getSSID() (string, error) {
	// Only works on Mac
	app0 := "/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport"
	app0_arg0 := "-I"
	cmd1 := exec.Command(app0, app0_arg0)

	app1 := "grep"
	app1_arg0 := "-w"
	app1_arg1 := "SSID"
	cmd2 := exec.Command(app1, app1_arg0, app1_arg1)

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	stdout, err := cmd2.StdoutPipe()
	cmd1.Start()

	if err != nil {
		fmt.Println(err)
	}
	if err := cmd2.Start(); err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)

	tokens := strings.Split(strings.TrimSpace(buf.String()), " ")

	if len(tokens) != 2 {
		return "", errors.New("Unable to extract SSID from airport resource")
	} else {
		return tokens[1], nil
	}
}

func runCalibration() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	ssid, err := getSSID()
	if err != nil {
		fmt.Println(err)
	}

	file_name := fmt.Sprintf("%s/.wif/speedtest_history/%s.csv", home, ssid)

	var include_header bool
	if _, err := os.Stat(file_name); err == nil {
		include_header = false
	} else {
		include_header = true
	}

	app := "speedtest"
	arg0 := "-f"
	arg1 := "csv"

	var cmd *exec.Cmd
	if include_header {
		arg2 := "--output-header"
		cmd = exec.Command(app, arg0, arg1, arg2)
	} else {
		cmd = exec.Command(app, arg0, arg1)
	}

	fmt.Println("Calibrating wif")
	stdout, cal_err := cmd.Output()
	if cal_err != nil {
		fmt.Println("speedtest not installed.")
		fmt.Println("Install speedtest here: https://www.speedtest.net/apps")
	}

	fmt.Println("Saving results to", file_name)
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Unable to open file")
	}
	defer f.Close()
	f.WriteString(string(stdout))

	return nil

}
