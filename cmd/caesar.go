package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	caesarDecode bool
	caesarShift  int
)

var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "(En/de)code using the Caesar cipher",
	Long:  `(En/de)code using the Caesar cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := caesarCipher(input, caesarShift, caesarDecode)
		fmt.Print(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(caesarCmd)
	caesarCmd.Flags().BoolVarP(&caesarDecode, "decode", "d", false, "Decode the input instead of encrypting")
	caesarCmd.Flags().IntVarP(&caesarShift, "shift", "s", 0, "Shift value for the Caesar cipher (required)")
	caesarCmd.MarkFlagRequired("shift") // Ensure shift is provided
}

// The Caesar cipher
func caesarCipher(input string, shift int, decode bool) string {
	if decode {
		shift = -shift // Reverse the shift for decodeion
	}
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r // Non-alphabet characters are returned as is
		}
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		// Shifting around the alphabet
		return base + (r-base+rune(shift)+26)%26
	}, input)
}
