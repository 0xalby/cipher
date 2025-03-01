package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	vigenereDecrypt bool
	vigenereKey     string
)

var vigenereCmd = &cobra.Command{
	Use:   "vigenere",
	Short: "Encrypt and decrypt using the Vigenère cipher",
	Long:  `Encrypt and decrypt using the Vigenère cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := vigenereCipher(input, vigenereKey, vigenereDecrypt)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(vigenereCmd)
	vigenereCmd.Flags().BoolVarP(&vigenereDecrypt, "decrypt", "d", false, "Decrypt the input instead of encrypting")
	vigenereCmd.Flags().StringVarP(&vigenereKey, "key", "k", "", "Key for the Vigenère cipher(required)")
	vigenereCmd.MarkFlagRequired("key")
}

// The Vigenère cipher
func vigenereCipher(input, key string, decrypt bool) string {
	key = strings.ToLower(key)
	keyIndex := 0 // Current key position
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}
		// The base value for lowercase or uppercase letters
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		// Converting the current key to a shift value
		keyChar := rune(key[keyIndex%len(key)]) - 'a'
		if decrypt {
			keyChar = -keyChar
		}
		// Going to the next key
		keyIndex++
		// Applying the shift and wrap around the alphabet
		return base + (r-base+keyChar+26)%26
	}, input)
}
