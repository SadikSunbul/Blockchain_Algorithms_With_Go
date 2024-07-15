package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Node represents a node in the Merkle Tree
type Node struct {
	Left  *Node
	Right *Node
	Data  []byte
}

// NewNode creates a new node with the given data
func NewNode(data []byte) *Node {
	return &Node{
		Data: data,
	}
}

// Hash computes the SHA-256 hash of the given data
func Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// BuildMerkleTree builds a Merkle Tree from the given data slices
func BuildMerkleTree(data [][]byte) *Node {
	var nodes []*Node

	// Create leaf nodes
	for _, d := range data {
		nodes = append(nodes, NewNode(Hash(d)))
	}

	// Build the tree from the leaf nodes
	for len(nodes) > 1 {
		var newLevel []*Node

		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				newNode := NewNode(Hash(append(nodes[i].Data, nodes[i+1].Data...)))
				newNode.Left = nodes[i]
				newNode.Right = nodes[i+1]
				newLevel = append(newLevel, newNode)
			} else {
				newLevel = append(newLevel, nodes[i])
			}
		}

		nodes = newLevel
	}

	return nodes[0]
}

// PrintTree prints the Merkle Tree in a readable format
func PrintTree(node *Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	fmt.Printf("%s", prefix)
	if isLast {
		fmt.Printf("└── ")
		prefix += "    "
	} else {
		fmt.Printf("├── ")
		prefix += "│   "
	}

	fmt.Printf("%s\n", hex.EncodeToString(node.Data))

	PrintTree(node.Left, prefix, node.Right == nil)
	PrintTree(node.Right, prefix, true)
}

func main() {
	data := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("merkle"),
		[]byte("tree"),
	}

	root := BuildMerkleTree(data)
	PrintTree(root, "", true)
}
