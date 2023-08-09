package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Utils struct {
	// Path to the entity
}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) GetGlobalPath() (string, error) {
	globalPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return "", err
	}
	globalPath = strings.Replace(globalPath, "\\", "/", -1)
	globalPath += "/"
	return globalPath, nil
}

func (u *Utils) GetPath(cmd *cobra.Command) (string, error) {
	globalPath, err := u.GetGlobalPath()
	if err != nil {
		return "", err
	}

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	path = strings.Replace(path, "\\", "/", -1)
	path = globalPath + path

	return path, nil
}

func (u *Utils) GetOutPath(cmd *cobra.Command) (string, error) {
	globalPath, err := u.GetGlobalPath()
	if err != nil {
		return "", err
	}

	outPath, err := cmd.Flags().GetString("out")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if outPath == "" {
		outPath = globalPath
	}

	return outPath, nil
}
