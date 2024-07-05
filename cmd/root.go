package foobar

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "1.0.0",
	Use:     "foo-bar",
	Short:   "foo-bar - a simple CLI to do random cool stuffs",
	Long: `stringer is a super fancy CLI written in Go using the Cobra CLI Framework
	
	Check out the sub commands to perform random fun things`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Woops. Something seems to have broken. '%s'", err)
		os.Exit(1)
	}
}
