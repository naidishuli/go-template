package command

import "github.com/spf13/cobra"

var LoadTempCmd = &cobra.Command{
    Use:   "loadTemp",
    Short: "Load temp data",
    Run:   loadTemp,
}

func loadTemp(c *cobra.Command, args []string) {

}
