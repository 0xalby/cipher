package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	caesarDecrypt bool
	caesarShift   int
)

var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Encrypt and decrypt using the Caesar cipher",
	Long:  `Encrypt and decrypt using the Caesar cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := caesarCipher(input, caesarShift, caesarDecrypt)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(caesarCmd)
	caesarCmd.Flags().BoolVarP(&caesarDecrypt, "decrypt", "d", false, "Decrypt the input")
	caesarCmd.Flags().IntVarP(&caesarShift, "shift", "s", 0, "Shift value for the Caesar cipher(required)")
	caesarCmd.MarkFlagRequired("shift")
}

// The Caesar cipher
func caesarCipher(input string, shift int, decrypt bool) string {
	if decrypt {
		shift = -shift
	}
	return strings.Map(func(r rune) rune {
		// Skipping non alphabet characters
		if !unicode.IsLetter(r) {
			return r
		}
		// Determining the base value for lowercase and uppercase letters
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		// Shifting around the alphabet
		return base + (r-base+rune(shift)+26)%26
	}, input)
}
