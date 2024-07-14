package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func main() {
	// ECC anahtarları oluştur
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Anahtar oluşturma hatası: %v", err)
	}

	// Ortak anahtarı al
	publicKey := &privateKey.PublicKey
	fmt.Printf("Public key:%x\n", publicKey)
	fmt.Printf("Private key:%x\n", privateKey)

	// Şifrelenecek metin
	message := []byte("2000 kelimeden oluşan metin...")

	// Metni şifrele
	encryptedMessage, err := encrypt(message, publicKey)
	if err != nil {
		log.Fatalf("Şifreleme hatası: %v", err)
	}

	fmt.Printf("Şifreli Metin: %s\n", hex.EncodeToString(encryptedMessage))

	// Şifreli metni çöz
	decryptedMessage, err := decrypt(encryptedMessage, privateKey)
	if err != nil {
		log.Fatalf("Çözme hatası: %v", err)
	}

	fmt.Printf("Çözülmüş Metin: %s\n", decryptedMessage)
}

// encrypt fonksiyonu, bir mesajı (message) belirtilen ortak anahtar (publicKey) kullanarak şifreler.
// Bu fonksiyon, mesajı ECIES (Elliptic Curve Integrated Encryption Scheme) kullanarak şifreler.
func encrypt(message []byte, publicKey *ecdsa.PublicKey) ([]byte, error) {
	// ECDSA ortak anahtarını ECIES ortak anahtarına dönüştürür.
	eciesPublicKey := ecies.ImportECDSAPublic(publicKey)
	fmt.Printf("eciesPublicKey:%x\n", eciesPublicKey)
	// Mesajı rastgele bir kaynak kullanarak (rand.Reader) şifreler ve şifreli mesajı döndürür.
	return ecies.Encrypt(rand.Reader, eciesPublicKey, message, nil, nil)
}

// decrypt fonksiyonu, şifreli bir mesajı (encryptedMessage) belirtilen özel anahtar (privateKey) kullanarak çözer.
// Bu fonksiyon, şifreli mesajı ECIES kullanarak çözer.
func decrypt(encryptedMessage []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// ECDSA özel anahtarını ECIES özel anahtarına dönüştürür.
	eciesPrivateKey := ecies.ImportECDSA(privateKey)
	fmt.Printf("eciesPrivateKey:%x\n", eciesPrivateKey)
	// Şifreli mesajı çözer ve orijinal mesajı döndürür.
	return eciesPrivateKey.Decrypt(encryptedMessage, nil, nil)
}
