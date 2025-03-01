package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	aesKey     string
	aesDecrypt bool
)

var aesCmd = &cobra.Command{
	Use:   "aes",
	Short: "Encrypt and decrypt using the AES cipher",
	Long:  `Encrypt and decrypt using the AES cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		if aesDecrypt {
			result, err := aesDecryptFunc(input, aesKey)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		} else {
			result, err := aesEncryptFunc(input, aesKey)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		}
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(aesCmd)
	aesCmd.Flags().BoolVarP(&aesDecrypt, "decrypt", "d", false, "Decrypt the input")
	aesCmd.Flags().StringVarP(&aesKey, "key", "k", "", "AES encryption/decryption key(32 bytes, required)")
	aesCmd.MarkFlagRequired("key")
}

// AES encryption
func aesEncryptFunc(input, key string) (string, error) {
	// Ensure the key is 32 bytes(AES-256)
	if len(key) != 32 {
		return "", fmt.Errorf("key must be exactly 32 bytes")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// Pad the input to be a multiple of the block size(AES block size is 16)
	paddedInput := pkcs7Pad([]byte(input), aes.BlockSize)
	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	// Create a new AES cipher mode(CBC mode)
	mode := cipher.NewCBCEncrypter(block, iv)
	// Encrypt the input
	encrypted := make([]byte, len(paddedInput))
	mode.CryptBlocks(encrypted, paddedInput)
	// Prepend the IV to the encrypted message for later use in decryption
	encrypted = append(iv, encrypted...)
	// Encode the encrypted result as a base64 string
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// AES decryption
func aesDecryptFunc(input, key string) (string, error) {
	// Ensure the key is 32 bytes(AES-256)
	if len(key) != 32 {
		return "", fmt.Errorf("key must be exactly 32 bytes")
	}
	// Decode the input from base64
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	// Get the IV from the first block(AES block size is 16)
	if len(decoded) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := decoded[:aes.BlockSize]
	encrypted := decoded[aes.BlockSize:]
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// Create a new AES cipher mode(CBC mode)
	mode := cipher.NewCBCDecrypter(block, iv)
	// Decrypt the ciphertext
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)
	// Remove padding
	decrypted, err = pkcs7Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// Adds PKCS7 padding
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Removes PKCS7 padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid data length for unpadding")
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > blockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	for i := len(data) - padding; i < len(data); i++ {
		if int(data[i]) != padding {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padding], nil
}
