package crypt

import (
	"crypto/aes"
	"encoding/base64"
)

// Helper function to repeat a byte multiple times
func bytesRepeating(b byte, count int) []byte {
	result := make([]byte, count)
	for i := range result {
		result[i] = b
	}
	return result
}

// Helper function to convert secret key into an AES-compatible key.
func convertToAESKey(secretKey []byte) []byte {
	// valid AES key sizes
	validKeySizes := []int{16, 24, 32}

	// Find the closest valid key size
	validSize := validKeySizes[0]
	for _, size := range validKeySizes {
		if len(secretKey) <= size {
			validSize = size
			break
		}
	}

	// Pad or truncate the secret key to the valid size
	aesKey := make([]byte, validSize)
	copy(aesKey, secretKey)

	return aesKey
}

// EncryptAES encrypts a plaintext using AES encryption in ECB mode
func EncryptAES(plainText, secretKey string) (string, error) {
	// Convert the secret key to a fixed-length AES key
	key := convertToAESKey([]byte(secretKey))

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad the plaintext to be a multiple of the block size
	plainTextBytes := []byte(plainText)
	padding := aes.BlockSize - len(plainTextBytes)%aes.BlockSize
	padText := append(plainTextBytes, bytesRepeating(byte(padding), padding)...)

	// Encrypt the padded plaintext using ECB mode
	ciphertext := make([]byte, len(padText))
	blockSize := block.BlockSize()
	for i := 0; i < len(padText); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], padText[i:i+blockSize])
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES decrypts an AES-encrypted string in ECB mode
func DecryptAES(ciphertext, secretKey string) (string, error) {
	// Convert the secret key to a fixed-length AES key
	key := convertToAESKey([]byte(secretKey))

	// Decode the base64-encoded ciphertext
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a buffer for the decrypted plaintext
	decryptedText := make([]byte, len(decodedCiphertext))
	blockSize := block.BlockSize()
	for i := 0; i < len(decodedCiphertext); i += blockSize {
		block.Decrypt(decryptedText[i:i+blockSize], decodedCiphertext[i:i+blockSize])
	}

	// Unpad the decrypted plaintext
	paddingByte := decryptedText[len(decryptedText)-1]
	padding := int(paddingByte)
	plainText := decryptedText[:len(decryptedText)-padding]

	return string(plainText), nil
}
