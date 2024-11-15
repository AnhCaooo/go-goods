// AnhCao 2024
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/AnhCaooo/go-goods/helpers"
)

// Read encryption key from config folder
// also having the trim empty space for key to ensure key are correct format
func ReadEncryptionKey(keyFilePath string) ([]byte, error) {
	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %s", err.Error())
	}

	return helpers.TrimSpaceForByte(key), nil
}

// Receives 2 paths (inputFile, outputFile) which are non-encrypted config file and encrypted config file.
// First, read non-encrypted data from given file path. Then do encrypt the data and write to desired file path
func EncryptFile(key []byte, decryptedFilePath, encryptedFilePath string) error {
	// Reading file
	plainText, err := os.ReadFile(decryptedFilePath)
	if err != nil {
		return fmt.Errorf("failed to decrypted read file: %s", err.Error())
	}

	// Encrypt data by receiving encryption key and plain data
	cipherText, err := encryptAES(key, plainText)
	if err != nil {
		return err
	}

	// Writing encrypted data to file path
	err = os.WriteFile(encryptedFilePath, cipherText, 0777)
	if err != nil {
		return fmt.Errorf("failed to write encrypted data to file: %s", err.Error())
	}
	return nil
}

// AES-GCM encryption
func encryptAES(key []byte, plainText []byte) ([]byte, error) {
	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create algorithm block: %s", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GCM mode: %s", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate random nonce: %s", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}

// Receives 2 paths (inputFile, outputFile) which are encrypted config file and non-encrypted config file.
// First, read encrypted data from given file path. Then do decrypt the data and write to desired file path
func DecryptFile(key []byte, encryptedFilePath, decryptedFilePath string) error {
	// Reading encrypted file
	cipherText, err := os.ReadFile(encryptedFilePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err.Error())
	}

	// Decrypt data by receiving encryption key and plain data
	plainText, err := decryptAES(key, cipherText)
	if err != nil {
		return err
	}

	// Writing decryption content
	err = os.WriteFile(decryptedFilePath, plainText, 0777)
	if err != nil {
		return fmt.Errorf("failed to write decrypted data to file: %s", err.Error())
	}
	return nil
}

// AES-GCM decryption
func decryptAES(key []byte, cipherText []byte) ([]byte, error) {
	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create block of algorithm: %s", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GCM mode: %s", err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, fmt.Errorf("cipherText too short")
	}

	// Detached nonce and decrypt
	nonce := cipherText[:nonceSize]
	cipherText = cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt file: %s", err.Error())
	}
	return plainText, nil
}
