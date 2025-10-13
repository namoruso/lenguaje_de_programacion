package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "taskcli",
		Short: "CLI sencillo para gestionar tareas (JSON en $HOME/.taskcli/tasks.json)",
		Long:  "taskcli es un CLI para gestionar tareas; soporta add, list, view, start, done, edit, rm.",
	}

	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdView)
	rootCmd.AddCommand(cmdStart)
	rootCmd.AddCommand(cmdDone)
	rootCmd.AddCommand(cmdEdit)
	rootCmd.AddCommand(cmdRemove)
	rootCmd.AddCommand(cmdVersion)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
