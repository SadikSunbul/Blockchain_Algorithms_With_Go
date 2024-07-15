package main

/*
Merkle Proofs (Merkle İspatları), bir Merkle Tree (Merkle Ağacı) yapısı kullanarak, belirli bir verinin bu ağaçta
bulunup bulunmadığını doğrulamak için kullanılan bir yöntemdir. Merkle Tree, özellikle blok zincirlerinde ve veri
bütünlüğünü doğrulamak için kullanılan bir veri yapısıdır.
*/
/*
Merkle Proofs Nasıl Çalışır?
Merkle Proofs, belirli bir verinin Merkle Tree'de bulunup bulunmadığını doğrulamak için kullanılır. İşlem şu şekildedir:

Veri Hash'i Hesaplanır: Doğrulanacak verinin hash değeri hesaplanır.

İspat Yolu (Proof Path) Oluşturulur: Verinin bulunduğu yaprak düğümden başlayarak, kök düğüme kadar olan yoldaki tüm
düğümlerin hash değerleri toplanır. Bu yol, ispat yolu (proof path) olarak adlandırılır.

Kök Hash'i Doğrulanır: İspat yolu kullanılarak, verinin hash değeriyle başlayarak, kök düğümün hash değeri hesaplanır.
Hesaplanan kök hash değeri, bilinen kök hash değeriyle karşılaştırılır. Eğer iki hash değeri eşleşiyorsa, veri Merkle
Tree'de bulunuyor demektir.
*/
func main() {

}
