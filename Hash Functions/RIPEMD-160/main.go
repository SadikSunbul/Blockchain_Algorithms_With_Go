package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	section1()
	/*
		RIPEMD-160:

		Bitcoin adreslerinin oluşturulmasında kullanılır.
	*/
}

func section1() {
	// Şifrelenecek veri
	data := []byte("hello world")

	// RIPEMD-160 hash nesnesi oluştur
	hash := ripemd160.New()

	// Veriyi hash nesnesine yaz
	hash.Write(data)

	// Hash değerini hesapla
	hashValue := hash.Sum(nil)

	// Hash değerini yazdır
	fmt.Printf("RIPEMD-160 hash: %x\n", hashValue)
}
