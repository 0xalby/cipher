package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	substitutionDecrypt  bool
	substitutionAlphabet string
)

var substitutionCmd = &cobra.Command{
	Use:   "substitution",
	Short: "Encrypt and decrypt using the Substitution cipher",
	Long:  `Encrypt and decrypt using the Substitution cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := substitutionCipher(input, substitutionAlphabet, substitutionDecrypt)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(substitutionCmd)
	substitutionCmd.Flags().BoolVarP(&substitutionDecrypt, "decrypt", "d", false, "Decrypt the input")
	substitutionCmd.Flags().StringVarP(&substitutionAlphabet, "alphabet", "a", "", "Substitution alphabet(26 unique letters and required)")
	substitutionCmd.MarkFlagRequired("alphabet")
}

// The Substitution cipher
func substitutionCipher(input, alphabet string, decrypt bool) string {
	if len(alphabet) != 26 {
		return "error alphabet must be exactly 26 letters long"
	}
	alphabet = strings.ToLower(alphabet)
	// The standard alphabet
	standardAlphabet := "abcdefghijklmnopqrstuvwxyz"
	if decrypt {
		// Swapping the standard alphabet and substitution alphabet for decryption
		standardAlphabet, alphabet = alphabet, standardAlphabet
	}
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			// A non alphabet characters gets returned as is
			return r
		}
		isUpper := unicode.IsUpper(r)
		r = unicode.ToLower(r)
		index := strings.IndexRune(standardAlphabet, r)
		if index == -1 {
			// Character not found
			return r
		}
		result := rune(alphabet[index])
		if isUpper {
			result = unicode.ToUpper(result)
		}
		return result
	}, input)
}
