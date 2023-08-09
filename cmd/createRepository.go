/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createRepositoryCmd represents the createRepository command
var createRepositoryCmd = &cobra.Command{
	Use:   "createRepository",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createRepository called")
	},
}

var Path string
var OutPath string

func init() {
	createRepositoryCmd.Flags().StringVarP(&Path, "path", "p", "", "Path to the entity")
	createRepositoryCmd.Flags().StringVarP(&OutPath, "out", "o", "", "Path to the output directory")
	rootCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRepositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
