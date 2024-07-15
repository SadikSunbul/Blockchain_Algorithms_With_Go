package main

import (
	"fmt"
	"github.com/SadikSunbul/Blockchain_Algorithms_With_Go/merkleTree"
)

func main() {
	datas := []string{"sadik", "ali", "ahmet", "mehmet", "deneme"}

	merkletree := merkleTree.NewNode()

	node := merkletree.MerkleTreeCreate(merkleTree.StringArrToNodeArr(datas))

	fmt.Println("Node read:")
	node.ReatTree("")
}
