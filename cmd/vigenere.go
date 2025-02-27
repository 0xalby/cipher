package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	vigenereDecode bool
	vigenereKey    string
)

var vigenereCmd = &cobra.Command{
	Use:   "vigenere",
	Short: "(En/de)code using the Vigenère cipher",
	Long:  `(En/de)code using the Vigenère cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := vigenereCipher(input, vigenereKey, vigenereDecode)
		fmt.Print(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(vigenereCmd)
	vigenereCmd.Flags().BoolVarP(&vigenereDecode, "decode", "d", false, "Decode the input instead of encrypting")
	vigenereCmd.Flags().StringVarP(&vigenereKey, "key", "k", "", "Key for the Vigenère cipher(required)")
	vigenereCmd.MarkFlagRequired("key")
}

// The Vigenère cipher
func vigenereCipher(input, key string, decode bool) string {
	key = strings.ToLower(key)
	keyIndex := 0
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		keyChar := rune(key[keyIndex%len(key)]) - 'a'
		if decode {
			keyChar = -keyChar
		}
		keyIndex++
		return base + (r-base+keyChar+26)%26
	}, input)
}
