package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	section2()
}

func section2() {
	//section 1 ile yanı işi yapar
	fmt.Printf("Hello World Hashi:%x\n", sha256.Sum256([]byte("Hello World")))
}

func section1() {
	sh := sha256.New() //yeni bir hash oluşturduk

	sh.Write([]byte("Hello World")) //veriyi hash'e yazdık

	fmt.Printf("%x\n", sh.Sum(nil)) //hash'i ekrana yazdık

	//sh.Reset() //bunu demez ısek altaki ile beraber ustekınıde ekelr yanı hasın cevabı aslında su olur "Hello WorldHello World"

	sh.Write([]byte("Hello World"))

	fmt.Printf("%x\n", sh.Sum(nil))
	sh.Reset()
	sh.Write([]byte("Hello WorldHello World"))
	fmt.Printf("%x\n", sh.Sum(nil))

	/* CIKTILAR:
	a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e
	60b3aaa7a9c94c59ff828b651205b1c5409eaf492505cd4a0a45133013b5a028
	60b3aaa7a9c94c59ff828b651205b1c5409eaf492505cd4a0a45133013b5a028
	*/
}

func section3() {
	// Dosya içeriğinin hash değerini hesaplama
	filePath := "example.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Dosya açılırken hata oluştu: %v", err)
	}
	defer file.Close()

	hashFile := sha256.New()
	if _, err := io.Copy(hashFile, file); err != nil {
		log.Fatalf("Dosya okunurken hata oluştu: %v", err)
	}

	hashFileValue := hashFile.Sum(nil)
	fmt.Printf("Dosya İçeriği SHA-256 Hash: %x\n", hashFileValue)
}
