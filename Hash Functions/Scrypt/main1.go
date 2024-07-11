package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
)

func main() {
	// Şifrelenecek veri
	password := []byte("mySecretPassword")

	// Rastgele tuz (salt) oluştur
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}

	// Scrypt parametreleri
	N := 16384   // CPU/memory maliyeti
	r := 8       // Blok boyutu
	p := 1       // Paralellik faktörü
	keyLen := 32 // İstenilen hash uzunluğu

	// Scrypt hash değeri hesapla
	hash, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	// Hash değerini ve tuzu yazdır
	fmt.Printf("Scrypt hash: %x\n", hash)
	fmt.Printf("Salt: %x\n", salt)

	// Doğrulama işlemi
	valid := verifyPassword(password, salt, hash, N, r, p, keyLen)
	fmt.Printf("Password validation: %t\n", valid)
}

func verifyPassword(password, salt, expectedHash []byte, N, r, p, keyLen int) bool {
	hash, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	return slowEquals(hash, expectedHash)
}

// slowEquals, zamanlama saldırılarına karşı koruma sağlamak için kullanılır
func slowEquals(a, b []byte) bool {
	diff := uint32(len(a)) ^ uint32(len(b))
	for i := 0; i < len(a) && i < len(b); i++ {
		diff |= uint32(a[i] ^ b[i])
	}
	return diff == 0
}
