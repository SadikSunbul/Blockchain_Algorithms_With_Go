package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"log"
)

func main() {
	section1()
	/*
		Equihash:

		Zcash gibi projelerde kullanılır.
	*/
}

const (
	N = 200
	K = 9
)

func section1() {
	// Rastgele bir nonce oluştur
	nonce := make([]byte, 32)
	_, err := rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}

	// Equihash çözümünü bul
	solution := solveEquihash(nonce)

	// Çözümü yazdır
	fmt.Printf("Equihash solution: %x\n", solution)
}

func solveEquihash(nonce []byte) []byte {
	// Blake2b hash fonksiyonu ile nonce'yi hashle
	hash := blake2b.Sum256(nonce)

	// Equihash çözümünü bulmak için basitleştirilmiş bir yöntem kullanıyoruz
	// Gerçek Equihash çözümü, N ve K parametrelerine dayalı karmaşık bir işlemdir
	// Burada sadece hash değerini kullanarak basit bir çözüm üretiyoruz

	// Çözümü byte dizisine dönüştür
	solution := make([]byte, 32)
	binary.LittleEndian.PutUint64(solution[:8], binary.LittleEndian.Uint64(hash[:8]))
	binary.LittleEndian.PutUint64(solution[8:16], binary.LittleEndian.Uint64(hash[8:16]))
	binary.LittleEndian.PutUint64(solution[16:24], binary.LittleEndian.Uint64(hash[16:24]))
	binary.LittleEndian.PutUint64(solution[24:32], binary.LittleEndian.Uint64(hash[24:32]))

	return solution
}
