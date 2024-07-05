package foobar

import (
	"fmt"
	"os"

	"github.com/alanjose10/foo-bar-cli/pkg/joke"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "joke",
	Short: "Tell me a random joke",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := joke.GetRandomJoke()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s :D\n", res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
