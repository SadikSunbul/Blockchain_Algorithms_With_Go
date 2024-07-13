package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {
	// RSA anahtarları oluştur
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Anahtar oluşturma hatası: %v", err)
	}

	// Ortak anahtarı al
	publicKey := &privateKey.PublicKey

	// Şifrelenecek mesaj
	message := []byte("Bu bir test mesajıdır.")

	// Mesajı şifrele
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		log.Fatalf("Şifreleme hatası: %v", err)
	}

	fmt.Printf("Şifreli Mesaj: %x\n", ciphertext)

	// Şifreli mesajı çöz
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		log.Fatalf("Şifre çözme hatası: %v", err)
	}

	fmt.Printf("Çözülmüş Mesaj: %s\n", plaintext)
}
