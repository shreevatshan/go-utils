package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/shreevatshan/go-utils/std"
)

// Helper function to repeat a byte multiple times
func bytesRepeating(b byte, count int) []byte {
	result := make([]byte, count)
	for i := range result {
		result[i] = b
	}
	return result
}

// Helper functions to convert secret key into an AES-128 compatible key.
func ConvertToAES128Key(secretKey []byte) []byte {
	// Pad or truncate the secret key to the valid size
	aesKey := make([]byte, 16)
	copy(aesKey, secretKey)

	return aesKey
}

// Helper functions to convert secret key into an AES-192 compatible key.
func ConvertToAES192Key(secretKey []byte) []byte {
	// Pad or truncate the secret key to the valid size
	aesKey := make([]byte, 24)
	copy(aesKey, secretKey)

	return aesKey
}

// Helper functions to convert secret key into an AES-256 compatible key.
func ConvertToAES256Key(secretKey []byte) []byte {
	// Pad or truncate the secret key to the valid size
	aesKey := make([]byte, 32)
	copy(aesKey, secretKey)

	return aesKey
}

func AESEncrypt(plainText, secretKey, iv []byte) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	if len(iv) != block.BlockSize() {
		return "", fmt.Errorf("iv size is incorrect")
	}

	stream := cipher.NewCBCEncrypter(block, iv)

	padding := aes.BlockSize - len(plainText)%aes.BlockSize
	plainText = append(plainText, bytesRepeating(byte(padding), padding)...)

	ciphertext := make([]byte, len(plainText))
	stream.CryptBlocks(ciphertext, plainText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AESDecrypt(encrypted, secretKey, iv []byte) (string, error) {

	ciphertext, err := std.DecodeBase64(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	if len(iv) != block.BlockSize() {
		return "", fmt.Errorf("iv size is incorrect")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	stream := cipher.NewCBCDecrypter(block, iv)

	plainText := make([]byte, len(ciphertext))
	stream.CryptBlocks(plainText, []byte(ciphertext))

	padding := int(plainText[len(plainText)-1])
	plainText = plainText[:len(plainText)-padding]

	return string(plainText), nil
}
