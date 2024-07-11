package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
)

func main() {
	section1()
	/*
		Scrypt:

		Litecoin gibi projelerde kullanılır.
	*/
}

func section1() {
	// Şifrelenecek veri
	password := []byte("mySecretPassword")

	// Rastgele tuz (salt) oluştur
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}

	// Scrypt parametreleri
	N := 16384   // CPU/memory maliyeti : 1'den büyük bir iki kuvveti olması gereken bir CPU/bellek maliyeti parametresidir.
	r := 8       // Blok boyutu
	p := 1       // Paralellik faktörü :  r ve p, r * p < 2³⁰'yi karşılamalıdır.
	keyLen := 32 // İstenilen hash uzunluğu

	// Scrypt hash değeri hesapla
	hash, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	// Hash değerini ve tuzu yazdır
	fmt.Printf("Scrypt hash: %x\n", hash)
	fmt.Printf("Salt: %x\n", salt)
}
