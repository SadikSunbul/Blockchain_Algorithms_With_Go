package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	section1()
	/*
		Keccak-256:

		Ethereum gibi projelerde kullanılır.
	*/
}

func section1() {
	/*
			Go dilinde Keccak-256 hash fonksiyonunu kullanmak için golang.org/x/crypto/sha3 paketini kullanabilirsiniz.
		İşte Keccak-256 hash fonksiyonunu kullanarak bir örnek kod:
	*/
	// Şifrelenecek veri
	data := []byte("hello world")

	// Keccak-256 hash nesnesi oluştur
	hash := sha3.New256()

	// Veriyi hash nesnesine yaz
	hash.Write(data)

	// Hash değerini hesapla
	hashValue := hash.Sum(nil)

	// Hash değerini yazdır
	fmt.Printf("Keccak-256 hash: %x\n", hashValue)
}
