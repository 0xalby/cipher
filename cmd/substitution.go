package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	substitutionAlphabet string
	substitutionDecode   bool
)

var substitutionCmd = &cobra.Command{
	Use:   "substitution",
	Short: "Encode and decode using the Substitution cipher",
	Long:  `Encode and decode using the Substitution cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := substitutionCipher(input, substitutionAlphabet, substitutionDecode)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(substitutionCmd)
	substitutionCmd.Flags().BoolVarP(&substitutionDecode, "decode", "d", false, "Decode the input")
	substitutionCmd.Flags().StringVarP(&substitutionAlphabet, "alphabet", "a", "", "Substitution alphabet(26 unique letters and required)")
	substitutionCmd.MarkFlagRequired("alphabet")
}

// Substitution cipher implementation
func substitutionCipher(input, alphabet string, decode bool) string {
	if len(alphabet) != 26 {
		return "error alphabet must be exactly 26 letters long"
	}
	alphabet = strings.ToLower(alphabet)
	standardAlphabet := "abcdefghijklmnopqrstuvwxyz"
	if decode {
		// Swapping the standard alphabet and substitution alphabet for decodeion
		standardAlphabet, alphabet = alphabet, standardAlphabet
	}
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r // Non-alphabet characters are returned as is
		}
		isUpper := unicode.IsUpper(r)
		r = unicode.ToLower(r)
		index := strings.IndexRune(standardAlphabet, r)
		if index == -1 {
			return r // Character not found in the alphabet
		}
		result := rune(alphabet[index])
		if isUpper {
			result = unicode.ToUpper(result)
		}
		return result
	}, input)
}
