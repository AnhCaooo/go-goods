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

// ReadEncryptionKey reads an encryption key from a file and ensures the key is correctly formatted.
//
// NOTE: keep encryption key in secret place and DO NOT PUBLIC its. Also this function does not store any key itself, just read and return.
//
// This function reads the contents of the specified file, trims any extra whitespace or
// newlines to ensure the key is in the correct format, and returns the sanitized key as a byte slice.
// If the file cannot be read, an error is returned.
//
// Example usage:
//
//	key, err := crypto.ReadEncryptionKey(keyFilePath)
//	if err != nil {
//	    return fmt.Errorf("error reading encryption key: %w", err)
//	}
//
// Parameters:
//   - keyFilePath: The path to the file containing the encryption key.
//
// Returns:
//   - []byte: The sanitized encryption key.
//   - error: An error if the file cannot be read or processed.
func ReadEncryptionKey(keyFilePath string) ([]byte, error) {
	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %s", err.Error())
	}

	return helpers.TrimSpaceForByte(key), nil
}

// EncryptFile reads data from a plaintext configuration file, encrypts it using a provided key,
// and writes the encrypted data to a specified output file.
//
// This function performs the following steps:
//  1. Reads the plaintext data from the file at `decryptedFilePath`.
//  2. Encrypts the data using the provided encryption key (`key`).
//  3. Writes the encrypted data to the file at `encryptedFilePath`.
//
// If any step fails, an error is returned.
//
// Example usage:
//
//	if err = crypto.DecryptFile(key, decryptedFilePath, encryptedFilePath); err != nil {
//		return err
//	}
//
// Parameters:
//   - key: The ENCRYPTION KEY used to encrypt the data.
//   - decryptedFilePath: The PATH to the plaintext configuration file (input).
//   - encryptedFilePath: The PATH to the encrypted configuration file (output).
//
// Returns:
//   - error: An ERROR if any step (reading, encrypting, or writing) fails.
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

// DecryptFile reads encrypted data from a file, decrypts it using a provided key,
// and writes the decrypted data to a specified output file.
//
// This function performs the following steps:
//  1. Reads the encrypted data from the file at `encryptedFilePath`.
//  2. Decrypts the data using the provided decryption key (`key`).
//  3. Writes the decrypted data to the file at `decryptedFilePath`.
//
// If any step fails, an error is returned.
//
// EXAMPLE USAGE:
//
//	if err = crypto.DecryptFile(key, encryptedFilePath, decryptedFilePath); err != nil {
//		return err
//	}
//
// PARAMETERS:
//   - key: The DECRYPTION KEY used to decrypt the data.
//   - encryptedFilePath: The PATH to the encrypted file (input).
//   - decryptedFilePath: The PATH to the decrypted file (output).
//
// RETURNS:
//   - error: An ERROR if any step (reading, decrypting, or writing) fails.
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
