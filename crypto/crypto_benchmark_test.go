package crypto

import (
	"crypto/rand"
	"testing"
)

func BenchmarkEncryptAES(b *testing.B) {
	key := make([]byte, 32) // AES-256
	rand.Read(key)

	data := make([]byte, 1024) // 1 KB payload
	rand.Read(data)

	b.ResetTimer()
	for b.Loop() {
		_, _ = encryptAES(key, data)
	}
}

func BenchmarkDecryptAES(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)

	data := make([]byte, 1024)
	rand.Read(data)

	cipherText, _ := encryptAES(key, data)

	b.ResetTimer()
	for b.Loop() {
		_, _ = decryptAES(key, cipherText)
	}
}
