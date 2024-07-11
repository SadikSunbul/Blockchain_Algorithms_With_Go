package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	section5()
}

func section1() {
	data := []byte("Hello World")
	sha3 := sha3.New256()
	sha3.Write(data)
	fmt.Printf("%x\n", sha3.Sum(nil))
}

func section2() {
	data := []byte("Hello World")
	// SHA3-512 hash hesaplama
	hash := sha3.New512()
	hash.Write(data)
	result := hash.Sum(nil)

	// Hash değerini yazdırma
	fmt.Printf("SHA3-512: %x\n", result)
}

func section3() {
	// SHA3-512 hash nesnesi oluştur
	sha := sha3.New512()

	// Veriyi hash nesnesine yaz
	sha.Write([]byte("Hello World"))

	// Hash değerini hesapla
	hash := sha.Sum(nil)

	// Hash değerini yazdır
	fmt.Printf("SHA3-512 hash: %x\n", hash)

	// Shake-256 hash değeri hesapla
	shakeHash := make([]byte, 32)                      // 32 byte (256 bit) çıktı uzunluğu
	sha3.ShakeSum256(shakeHash, []byte("Hello World")) //istediğiniz boyutta şifreleme yapmayı sasğlar tek yapmanız gereken shakeHash buyuklugunu artırıp dsurmek

	// Shake-256 hash değerini yazdır
	fmt.Printf("Shake-256 hash: %x\n", shakeHash)
}

func section4() {
	fmt.Printf("Hash:%x\n", sha3.Sum256([]byte("Hello World")))
	//hızlı kullanımı
}

func section5() {
	out := make([]byte, 32)
	msg := []byte("The quick brown fox jumps over the lazy dog")

	//özelelştirilebilir
	c1 := sha3.NewCShake256([]byte("NAME"), []byte("Partition1"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 2: Farklı özelleştirme dizesi farklı özet üretir
	c1 = sha3.NewCShake256([]byte("NAME"), []byte("Partition2"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 3:Daha uzun çıktı uzunluğu daha uzun özet üretir
	out = make([]byte, 64)
	c1 = sha3.NewCShake256([]byte("NAME"), []byte("Partition1"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 4: Sonraki okuma farklı sonuç üretir
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))
	/*
			c1.Read(out) işlemi ilk kez çağrıldığında, out değişkenine 64 byte uzunluğunda bir hash değeri yazılır.
		Ardından, c1.Read(out) işlemi tekrar çağrıldığında, cSHAKE-256
		nesnesi içindeki veri tüketilmiş olduğu için farklı bir hash de
		ğeri üretilir. Bu, cSHAKE-256 algoritmasının çalışma şekliyle
		ilgilidir ve her Read işlemi, önceki Read işlemlerinden farklı bir
		sonuç üretir.
	*/
	/*
		Çıktı:

		a90a4c6ca9af2156eba43dc8398279e6b60dcd56fb21837afe6c308fd4ceb05b
		a8db03e71f3e4da5c4eee9d28333cdd355f51cef3c567e59be5beb4ecdbb28f0
		a90a4c6ca9af2156eba43dc8398279e6b60dcd56fb21837afe6c308fd4ceb05b9dd98c6ee866ca7dc5a39d53e960f400bcd5a19c8a2d6ec6459f63696543a0d8
		85e73a72228d08b46515553ca3a29d47df3047e5d84b12d6c2c63e579f4fd1105716b7838e92e981863907f434bfd4443c9e56ea09da998d2f9b47db71988109
	*/
}
