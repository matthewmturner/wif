/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// Source: https://www.tomsguide.com/features/5g-vs-4g
// Download speeds in Mbps, prefixed with "DOWN_"
const DOWN_5G = 70_000_000
const DOWN_4G = 40_000_000

// 5G Source: https://9to5mac.com/2021/07/13/t-mobile-leads-5g-average-speeds-close-100-mbps/
// 4G Source: https://www.verizon.com/articles/4g-lte-speeds-vs-your-home-network/
// Upload speeds in Mbps, prefiex with "UP_"
const UP_5G = 10_000_000
const UP_4G = 3_000_000

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Get file information and expected up or download speed",
	Long: `Analyze is a tool to help reason about network and application performance (in the context of networks).  
	It provides file information from the API's perspective (i.e. the 'Content-type'), and back of the envelop 
	expected upload and download times.
	
File name - provided by user
File type - derived from net/http 'DetectContentType'
File size - from 'fileStat'
File last modified - from 'fileStat'
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("analyze accepts exactly 1 file path argument")
			return
		}
		file := args[0]
		fileStat, err := os.Stat(file)
		if err != nil {
			fmt.Printf("Unable to open file: %s\n", file)
			return
		}
		if fileStat.IsDir() {
			fmt.Printf("Analyze only accepts file paths, not directories")
			return
		}

		printAnalyzeResults(fileStat, file)

	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func formatFileSize(size int64) string {

	if size < 1_000 {
		return fmt.Sprintf("%d bytes", size)
	} else if size < 1_000_000 {
		return fmt.Sprintf("%f kilobytes", float32(size)/1_000)
	} else if size < 1_000_000_000 {
		return fmt.Sprintf("%f megabytes", float32(size)/1_000_000)
	} else {
		return "Too big"
	}
}

func getFileType(path string) string {

	// Open File
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	contentType, err := getFileContentType(f)
	if err != nil {
		panic(err)
	}
	return contentType
}

func getFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func printFileInfo(fileStat fs.FileInfo, file string) {
	size := fileStat.Size()

	// File information
	fmt.Println("\nFile Information")
	fmt.Println("File Name:", fileStat.Name())                                             // Base name of the file
	fmt.Println("File Type:", getFileType(file))                                           // File type
	fmt.Println("Size:", formatFileSize(size))                                             // Length in bytes for regular files
	fmt.Println("Last Modified:", fileStat.ModTime().Format("January 2, 2006 3:04:05 PM")) // Last modification time
}

func printPerformanceResults(size int64) {
	// Expected performance
	fmt.Println("\nExpected Network Speed Information")
	fmt.Println("5G Download: ", float32(size)/(DOWN_5G/8), "seconds")
	fmt.Println("4G Download: ", float32(size)/(DOWN_4G/8), "seconds")
	fmt.Println("5G Upload: ", float32(size)/(UP_5G/8), "seconds")
	fmt.Println("4G Upload: ", float32(size)/(UP_4G/8), "seconds")
}

func printAnalyzeResults(fileStat fs.FileInfo, file string) {

	size := fileStat.Size()

	printFileInfo(fileStat, file)
	printPerformanceResults(size)

}
