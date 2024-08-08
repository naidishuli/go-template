package main

import (
    "fmt"
    "os"

    "bets/cmd/command"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "cmd",
    Short: "betstion entry points, executables and more",
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
