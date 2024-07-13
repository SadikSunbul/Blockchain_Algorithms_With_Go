package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

/*
AES (Advanced Encryption Standard):

AES, 128, 192 ve 256 bitlik anahtar boyutları ile çalışır ve yüksek güvenlik seviyeleri sunar. Blockchain
teknolojisinde, veri şifreleme ve şifre çözme işlemleri için yaygın olarak kullanılır.
*/

// Şifreleme işlemini gerçekleştiren fonksiyon
func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	// NewCipher, AES şifreleme algoritması için yeni bir cipher.Block oluşturur.
	// Anahtar (key) parametresi, AES-128, AES-192 veya AES-256'yı seçmek için 16, 24 veya 32 baytlık bir AES anahtarı olmalıdır.
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// aes.BlockSize, AES algoritmasının blok boyutunu (16 byte) döndürür.
	// ciphertext dizisi, IV (Initialization Vector) ve şifrelenmiş veriyi tutacak şekilde oluşturulur.

	iv := ciphertext[:aes.BlockSize]
	// IV, şifrelenmiş verinin başında bulunan 16 byte'lık bir dizidir.
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		// io.ReadFull, rastgele bir IV oluşturmak için kullanılır.
		// rand.Reader, rastgele veri üreten bir io.Reader'dır.
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	// cipher.NewCFBEncrypter, CFB (Cipher Feedback) modunda bir şifreleme akışı oluşturur.
	// CFB modu, veriyi bloklara bölmeden akış şifrelemesi sağlar.
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// XORKeyStream, plaintext verisini şifrelemek için kullanılır.
	// ciphertext dizisinin IV'den sonraki kısmına şifrelenmiş veri yazılır.

	return ciphertext, nil
}

// Şifre çözme işlemini gerçekleştiren fonksiyon
func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	// NewCipher, AES şifreleme algoritması için yeni bir cipher.Block oluşturur.
	// Anahtar (key) parametresi, AES-128, AES-192 veya AES-256'yı seçmek için 16, 24 veya 32 baytlık bir AES anahtarı olmalıdır.
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("şifreli metin çok kısa")
	}

	iv := ciphertext[:aes.BlockSize]
	// IV, şifrelenmiş verinin başında bulunan 16 byte'lık bir dizidir.
	ciphertext = ciphertext[aes.BlockSize:]

	decryptedText := make([]byte, len(ciphertext))
	// Şifresi çözülmüş veriyi tutacak byte dizisi oluşturulur.
	stream := cipher.NewCFBDecrypter(block, iv)
	// cipher.NewCFBDecrypter, CFB modunda bir şifre çözme akışı oluşturur.
	stream.XORKeyStream(decryptedText, ciphertext)
	// XORKeyStream, şifrelenmiş veriyi çözmek için kullanılır.
	// decryptedText dizisine şifresi çözülmüş veri yazılır.

	return decryptedText, nil
}

func main() {
	// Şifreleme için kullanılacak anahtar
	key := []byte("0123456789abcdef0123456789abcdef") // 32 byte (256 bit)

	// Şifrelenecek veri
	plaintext := []byte("Merhaba, dünya!")

	// Şifreleme işlemi
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		panic(err)
	}

	// Şifrelenmiş veriyi hex formatında yazdır
	fmt.Printf("Ciphertext: %s\n", hex.EncodeToString(ciphertext))
	// hex.EncodeToString, ciphertext dizisini hexadecimal bir dizeye dönüştürür ve yazdırır.

	// Şifre çözme işlemi
	decryptedText, err := decrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}

	// Şifresi çözülmüş veriyi yazdır
	fmt.Printf("Decrypted text: %s\n", decryptedText)
	// Şifresi çözülmüş veriyi yazdırır.
}
