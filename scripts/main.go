package main

import (
    "fmt"
    "os"

    "go-template/scripts/command"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "scripts",
    Short: "General scripts for my up, seeders, data loaders etc",
}

func init() {
    rootCmd.AddCommand(command.LoadTempCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
