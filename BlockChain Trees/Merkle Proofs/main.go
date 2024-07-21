package main

import (
	"fmt"
	"github.com/SadikSunbul/Blockchain_Algorithms_With_Go/merkleTree"
)

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
	veriBloklari := []string{
		"blok1",
		"blok2",
		"blok3",
		"blok4",
		"blok5",
		"blok6",
	}

	merkle := merkleTree.NewNode()
	root := merkle.MerkleTreeCreate(merkleTree.StringArrToNodeArr(veriBloklari))
	Rooot = *root
	hashData := merkleTree.Hash([]byte("blok2"))

	Serch(root, hashData)

	fmt.Printf("root: %v\n", root)
}

func isItleaf(n *merkleTree.Node) bool {
	if n == nil {
		return false
	}
	if n.Left == nil && n.Right == nil {
		return true
	}
	return false
}

func Serch(root *merkleTree.Node, hashData [32]byte) bool {
	path := []*merkleTree.Node{}

	var rootre merkleTree.Node = *root
	path, _ = rekursif(&rootre, hashData, path)
	path = path[:len(path)-1]

	fmt.Printf("proof path: %x\n", path)
	return true
}

var Rooot merkleTree.Node

//func rekursif(root *merkleTree.Node, hashData [32]byte, path []*merkleTree.Node) []*merkleTree.Node {
//	if root == nil {
//		return path
//	}
//
//	if isItleaf(root) {
//		if hashData == root.Data {
//			fmt.Println("merkleTree'de bulundu")
//			path = append(path, root)
//			return path
//		}
//		path = path[:len(path)-1]
//	}
//	path = append(path, root)
//
//	path = rekursif(root.Left, hashData, path)
//	path = rekursif(root.Right, hashData, path)
//
//	return path
//}

func rekursif(root *merkleTree.Node, hashData [32]byte, path []*merkleTree.Node) ([]*merkleTree.Node, bool) {
	if root == nil {
		return path, false
	}
	path = append(path, root)
	if isItleaf(root) {
		if hashData == root.Data {
			fmt.Println("merkleTree'de bulundu")
			path = append(path, root)
			return path, true
		}
		path = path[:len(path)-1]
		return path, false
	}

	var found bool
	path, found = rekursif(root.Left, hashData, path)
	if found {
		//if path[1].Data == Rooot.Left.Data {
		//	path = path[:len(path)-1]
		//}
		return path, true
	}
	if len(path) >= 2 && path[1].Data == Rooot.Left.Data {
		path = path[:len(path)-1]
	}
	path, found = rekursif(root.Right, hashData, path)
	if found {
		return path, true
	}

	return path, false
}
