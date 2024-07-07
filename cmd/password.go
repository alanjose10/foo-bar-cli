package foobar

import (
	"errors"
	"fmt"
	"os"

	"github.com/alanjose10/foo-bar-cli/pkg/password"
	"github.com/spf13/cobra"
)

var length int
var specialChar bool
var upperCase bool
var numbers bool

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate passwords based on the provided rules",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateLength(length); err != nil {
			fmt.Printf("%d is too less for password length\n", length)
			os.Exit(1)
		}

		pc := password.NewPasswordConfig(
			password.WithLength(length),
			password.WithIncludeSpecialChar(specialChar),
			password.WithIncludeUpperCase(upperCase),
			password.WithIncludeNumbers(numbers),
		)
		fmt.Println(password.GeneratePassword(&pc))
	},
}

func validateLength(n int) error {
	if n < 10 {
		return errors.New("length must be greater than 10")
	} else if n >= 1000 {
		return errors.New("length must be less than 1000")
	}
	return nil
}

func init() {
	passwordCmd.Flags().IntVarP(&length, "length", "l", 20, "Password length")
	passwordCmd.Flags().BoolVarP(&specialChar, "inclide-special-char", "s", false, "Include special characters in password")
	passwordCmd.Flags().BoolVarP(&upperCase, "inclide-upper-case", "U", false, "Include upper case characters in password")
	passwordCmd.Flags().BoolVarP(&numbers, "inclide-numbers", "D", false, "Include numbers in password")
	rootCmd.AddCommand(passwordCmd)
}
