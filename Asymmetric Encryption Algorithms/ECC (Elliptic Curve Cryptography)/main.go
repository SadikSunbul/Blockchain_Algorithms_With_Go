package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

/*
ECC (Elliptic Curve Cryptography):

ECC, eliptik eğri matematiğini kullanarak güvenli anahtar üretimi, şifreleme ve şifre çözme işlemleri gerçekleştirir.
Daha küçük anahtar boyutları ile benzer güvenlik seviyeleri sunar. Blockchain teknolojisinde, düşük kaynaklı cihazlarda
ve yüksek güvenlik gereksinimleri olan uygulamalarda kullanılır.
*/

func main() {
	// Eliptik eğri parametreleri oluştur
	curve := elliptic.P256()

	// ECC anahtarları oluştur
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatalf("Anahtar oluşturma hatası: %v", err)
	}

	// Ortak anahtarı al
	publicKey := &privateKey.PublicKey

	// İmzalanacak mesaj
	message := []byte("Bu bir test mesajıdır.")

	// Mesajın SHA-256 özeti
	hashed := sha256.Sum256(message)

	// Mesajı imzala
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed[:])
	if err != nil {
		log.Fatalf("İmzalama hatası: %v", err)
	}

	// İmzayı doğrula
	if ecdsa.Verify(publicKey, hashed[:], r, s) {
		fmt.Println("İmza doğrulandı.")
	} else {
		fmt.Println("İmza doğrulanamadı.")
	}
}
