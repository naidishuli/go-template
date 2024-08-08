package main

import (
    "fmt"
    "os"

    "go-template/cmd/command"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "cmd",
    Short: "go-templatetion entry points, executables and more",
}

func init() {
    rootCmd.AddCommand(command.RunMigrationCmd)
    rootCmd.AddCommand(command.RunServerCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
