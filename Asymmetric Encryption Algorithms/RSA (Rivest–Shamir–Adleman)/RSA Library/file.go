package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Dosya oluşturma hatası: %v", err)
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		log.Fatalf("PEM kodlama hatası: %v", err)
	}
}

func loadPEMKey(fileName string) *rsa.PrivateKey {
	keyData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Dosya okuma hatası: %v", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		log.Fatalf("PEM bloğu bulunamadı")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Anahtar çözme hatası: %v", err)
	}

	return privateKey
}

func main() {
	// RSA anahtarları oluştur
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Anahtar oluşturma hatası: %v", err)
	}

	// Anahtarı dosyaya kaydet
	savePEMKey("private.pem", privateKey)

	// Anahtarı dosyadan yükle
	loadedPrivateKey := loadPEMKey("private.pem")

	// Ortak anahtarı al
	publicKey := &loadedPrivateKey.PublicKey

	// Şifrelenecek mesaj
	message := []byte("Bu bir test mesajıdır.")

	// Mesajı şifrele
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		log.Fatalf("Şifreleme hatası: %v", err)
	}

	fmt.Printf("Şifreli Mesaj: %x\n", ciphertext)

	// Şifreli mesajı çöz
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, loadedPrivateKey, ciphertext)
	if err != nil {
		log.Fatalf("Şifre çözme hatası: %v", err)
	}

	fmt.Printf("Çözülmüş Mesaj: %s\n", plaintext)
}
