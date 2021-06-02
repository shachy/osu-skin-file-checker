/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func SkfContainsName(n []string) []string {
	var missingFiles []string
	var found bool
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	exPath := strings.ReplaceAll(filepath.Dir(ex), "\\", "/") + "/donotmove.txt"
	doc, err := os.Open(exPath)
	if err != nil {
		fmt.Println(err)
	}
	defer doc.Close()
	scanner := bufio.NewScanner(doc)
	for scanner.Scan() {
		found = false
		for _, file := range n {
			if strings.Contains(scanner.Text(), file) {
				found = true
				break
			}
		}
		if found {
			continue
		} else {
			missingFiles = append(missingFiles, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return missingFiles
}

func CheckFolder(args []string) error {
	var DIRPATH = args[0]
	file, err := os.Open(DIRPATH)
	if err != nil {
		return err
	}
	defer file.Close()
	files, err := file.Readdirnames(0)
	if err != nil {
		return err
	}

	msFiles := SkfContainsName(files)

	for i := 0; i < len(msFiles); i++ {
		fgGreen := color.New(color.FgHiGreen).PrintfFunc()
		fmt.Print("Missing file [")
		fgGreen("%v", i+1)
		fmt.Printf("] -> %v\n", msFiles[i])
	}
	return nil
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks a skin folder",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := CheckFolder(args); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
