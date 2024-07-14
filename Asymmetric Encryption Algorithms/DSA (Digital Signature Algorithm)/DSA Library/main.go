package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	// DSA parametreleri oluştur
	params := new(dsa.Parameters)
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		log.Fatalf("Parametre oluşturma hatası: %v", err)
	}

	// DSA anahtarları oluştur
	privateKey := new(dsa.PrivateKey)
	privateKey.PublicKey.Parameters = *params
	if err := dsa.GenerateKey(privateKey, rand.Reader); err != nil {
		log.Fatalf("Anahtar oluşturma hatası: %v", err)
	}

	// Ortak anahtarı al
	publicKey := privateKey.PublicKey

	// İmzalanacak mesaj
	message := []byte("Bu bir test mesajıdır.")

	// Mesajın SHA-256 özeti
	hashed := sha256.Sum256(message)

	fmt.Println("public:", publicKey)
	fmt.Println("private:", privateKey)

	// Mesajı imzala
	r, s, err := dsa.Sign(rand.Reader, privateKey, hashed[:])
	if err != nil {
		log.Fatalf("İmzalama hatası: %v", err)
	}

	// İmzayı doğrula
	if dsa.Verify(&publicKey, hashed[:], r, s) {
		fmt.Println("İmza doğrulandı.")
	} else {
		fmt.Println("İmza doğrulanamadı.")
	}
}

/*
DSA Parametreleri Oluşturma:

dsa.GenerateParameters fonksiyonu kullanılarak DSA parametreleri oluşturulur. Bu parametreler, DSA algoritmasının çalışması için gereklidir.

DSA Anahtarları Oluşturma:

dsa.GenerateKey fonksiyonu kullanılarak özel ve ortak anahtarlar oluşturulur.

Mesajın Özeti Hesaplanması:

sha256.Sum256 fonksiyonu kullanılarak mesajın SHA-256 özeti hesaplanır.

Mesajın İmzalanması:

dsa.Sign fonksiyonu kullanılarak mesajın özeti imzalanır ve imza değerleri r ve s olarak elde edilir.

İmzanın Doğrulanması:

dsa.Verify fonksiyonu kullanılarak imza doğrulanır. Eğer imza doğru ise, mesajın orijinalliği ve bütünlüğü onaylanmış olur.
*/
