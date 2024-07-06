package foobar

import (
	"fmt"
	"os"

	"github.com/alanjose10/foo-bar-cli/pkg/ip"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "my-ip",
	Short: "Tell me my ip address",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := ip.GetMyIp()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", res)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
