package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"
)

/*
RSA (Rivest–Shamir–Adleman):

RSA, asimetrik şifreleme algoritmasıdır ve genellikle dijital imzalar, anahtar değişimi ve şifreleme işlemleri için
kullanılır. Blockchain teknolojisinde, işlemleri doğrulamak ve güvenli iletişim sağlamak için yaygın olarak kullanılır.
*/
func main() {
	rand.Seed(time.Now().UnixNano())
	p, q := KeyCreate()
	orjinalMesaj := 65

	sifreliMesaj, e, err := Encrypt(p, q, orjinalMesaj)
	if err != nil {
		log.Fatal("Şifreleme hatası!!!")
	}
	fmt.Println("Şifreli metin::", sifreliMesaj)

	cozulmusMessaj, err := Decrypt(sifreliMesaj, p, q, e)
	if err != nil {
		log.Fatal("Çözme hatası!!!")
	}
	fmt.Println("Orjinal metin:", cozulmusMessaj)
}

func KeyCreate() (p, q int) {
	//2 asalsayı üret
	// p ve q rastgele farklı asal sayıdır

	p = 61
	q = 53
	return
}

func Encrypt(p, q, message int) (int, int, error) {
	n := p * q
	//totien degeri
	phi := (p - 1) * (q - 1)
	e := primeNumberAmongThem(phi)
	if e == -1 {
		return 0, 0, errors.New("asal sayı bulunamadı")
	}
	//şifreleme işlemi
	// OrtakAnahtar(n,e)   m mesajı sifreleyelim c = m^e mod n
	// özelAnahatar(n,d)   c sifreli m = c^d mod n

	fmt.Println("n:", n, "e:", e, "phi:", phi)
	// Şifreleme işlemi: c = m^e mod n
	m := big.NewInt(int64(message))
	eBig := big.NewInt(int64(e))
	nBig := big.NewInt(int64(n))
	c := new(big.Int).Exp(m, eBig, nBig) // exp=c = m^e mod n işlemini yapar

	return int(c.Int64()), e, nil
}

func Decrypt(message, p, q, e int) (int, error) {
	n := p * q
	//totien degeri
	phi := (p - 1) * (q - 1)

	d, err := modInverse(e, phi)
	if err != nil {
		return 0, err
	}
	if d == -1 {
		return 0, errors.New("d değeri hesaplanamadı")
	}
	fmt.Println("d:", d, "e:", e, "phi:", phi)

	//şifreleme işlemi
	// OrtakAnahtar(n,e)   m mesajı sifreleyelim c = m^e mod n
	// özelAnahatar(n,d)   c sifreli m = c^d mod n

	// Şifre çözme işlemi: m = c^d mod n
	c := big.NewInt(int64(message))
	dBig := big.NewInt(int64(d))
	nBig := big.NewInt(int64(n))
	m := new(big.Int).Exp(c, dBig, nBig) // Exp :  m = c^d mod n

	return int(m.Int64()), nil
}

// extendedGCD fonksiyonu genişletilmiş Öklid algoritmasını uygular
func extendedGCD(a, b int) (g int, x int, y int) {
	if a == 0 {
		return b, 0, 1
	}
	g1, x1, y1 := extendedGCD(b%a, a)
	return g1, y1 - (b/a)*x1, x1
}

// modInverse fonksiyonu e'nin mod φ(n)'deki çarpmaya göre tersini hesaplar
func modInverse(e, phi int) (int, error) {
	g, x, _ := extendedGCD(e, phi)
	if g != 1 { //bunun 1 olmasını sundan istiyoruz e ve phi aralarında asal olması lazım aralarında asal ıseler 1 olmasılazım bu degerın
		return 0, fmt.Errorf("modüler ters mevcut değil\n")
	}
	return (x%phi + phi) % phi, nil //x yanı aradıgımız d negatıf olabılrı su durumda ama bız negatif olmasını ıstemeyoruz o yuzden pozitif halıne donduruz bu degeri
}

func primeNumberAmongThem(num int) int {
	primeFactors := make([]int, 0)
	for i := 2; i < num; i++ {
		if gcd(i, num) == 1 { // i ve num aralarında asal olması yeterli kendısını asal olmasına gerek yok
			primeFactors = append(primeFactors, i)
		}
	}
	if len(primeFactors) == 0 {
		return -1
	}

	return primeFactors[rand.Intn(len(primeFactors))]
}

// gcd fonksiyonu iki sayının en büyük ortak bölenini hesaplar
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func IsPrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
