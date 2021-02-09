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
	"errors"

	"github.com/ardtieboy/starlu/imageprocessing"
	"github.com/spf13/cobra"
)

// teaserCmd represents the teaser command
var teaserCmd = &cobra.Command{
	Use:   "teaser image_path",
	Short: "Crops, resizes and puts border on provided image in order to create a teaser",
	Long:  "Crops, resizes and puts border on provided image in order to create a teaser",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one or more filenames")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		println("Starting now... ☕️")
		for _, s := range args {
			// println("Processing " + s + " now...")
			output, err := imageprocessing.Crop(s)
			if err != nil {
				return err
			}
			println("➡️ " + output)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(teaserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teaserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teaserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
