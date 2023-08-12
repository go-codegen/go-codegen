package cmd

import (
	"fmt"
	"github.com/go-codegen/go-codegen/internal/colorPrint"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

type Utils struct {
	// Path to the entity
}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) GetPath(cmd *cobra.Command) (string, error) {
	globalPath, err := u.GetGlobalPath()
	if err != nil {
		return "", err
	}

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		colorPrint.PrintError(err)
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
		colorPrint.PrintError(err)
	}

	if outPath == "" {
		outPath = globalPath
	}

	return outPath, nil
}

func (u *Utils) ShowLoadingAnimation(done chan bool) {
	fmt.Print("   Waiting... ")
	spinners := []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}

	for i := 0; ; i++ {
		select {
		case <-done:
			u.clearLine()
			return
		default:
			fmt.Printf("\r%s", spinners[i%len(spinners)])
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (u *Utils) clearLine() {
	fmt.Printf("\r\033[K")
}
