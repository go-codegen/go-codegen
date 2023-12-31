/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/go-codegen/go-codegen/internal/colorPrint"
	repository_module "github.com/go-codegen/go-codegen/internal/modules/repository"
	"github.com/go-codegen/go-codegen/internal/parse"
	"github.com/go-codegen/go-codegen/internal/repository"
	"github.com/spf13/cobra"
)

// createRepositoryCmd represents the createRepository command
var createRepositoryCmd = &cobra.Command{
	Use:   "createRepository",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
 
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed xz
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils := NewUtils()

		globalPath, err := cmd.Flags().GetString("path")
		if err != nil {
			colorPrint.PrintError(err)
			return
		}

		path, err := utils.GetPath(cmd)
		if err != nil {
			colorPrint.PrintError(err)
			return
		}

		outPath, err := utils.GetOutPath(cmd)
		if err != nil {
			outPath = globalPath
		}

		switch args[0] {
		case "gorm":
			done := make(chan bool)
			go utils.ShowLoadingAnimation(done)

			module := repository_module.NewGorm()

			parseStruct := parse.NewStruct()

			structs, err := parseStruct.ParseStructInFiles(path)
			if err != nil {
				colorPrint.PrintError(err)
				return
			}

			body := repository.NewRepository(module, structs)
			err = body.Create(outPath)
			if err != nil {
				colorPrint.PrintError(err)
				return
			}

			done <- true

			colorPrint.PrintSuccess("Gorm repository created successfully")
		default:
			colorPrint.PrintError(fmt.Errorf("unknown module"))
		}

	},
}

func init() {

	createRepositoryCmd.Flags().StringP("path", "p", "", "Path to the entity")
	createRepositoryCmd.Flags().StringP("out", "o", "", "Path to the output directory")
	rootCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRepositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
