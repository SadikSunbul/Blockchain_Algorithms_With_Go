package merkleTree

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left       *Node
	Right      *Node
	Data       [32]byte
	DataString string //bunu sadece verıyı gormek ıcın yaptık
}

func NewNode() *Node {
	return &Node{}
}

func StringToByte(s string) []byte {
	return []byte(s)
}

func ByteToString(b []byte) string {
	return string(b)
}

func Hash(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func StringArrToNodeArr(datas []string) []*Node {
	var nodes []*Node
	for _, data := range datas {
		nodes = append(nodes, &Node{Data: Hash(StringToByte(data)), Left: nil, Right: nil, DataString: data})
	}
	return nodes
}

func (n *Node) MerkleTreeCreate(datas []*Node) *Node {
	return rekursifMerkelAdd(datas)
}

func rekursifMerkelAdd(datas []*Node) *Node {
	maxlen := len(datas)
	rekursifdata := make([]*Node, 0)
	for i := 0; i < len(datas); {
		if i+1 == maxlen {
			rekursifdata = append(rekursifdata, merkleTreeAdd(datas[i], nil))
		} else {
			rekursifdata = append(rekursifdata, merkleTreeAdd(datas[i], datas[i+1]))
		}
		i += 2
	}
	if len(rekursifdata) != 1 {
		return rekursifMerkelAdd(rekursifdata)
	}
	return rekursifdata[0]
}

func merkleTreeAdd(data1, data2 *Node) *Node {
	if data2 == nil {
		data2 = data1
	}
	return &Node{Left: data1, Right: data2, Data: Hash(bytes.Join([][]byte{data1.Data[:], data2.Data[:]}, []byte{}))}
}

//func (n *Node) ReatTree() {
//	if n == nil {
//		return
//	}
//	fmt.Printf("%x-%s\n", n.Data[:], n.DataString)
//	n.Left.ReatTree()
//	n.Right.ReatTree()
//}

func (n *Node) ReatTree(prefix string) {
	if n == nil {
		return
	}

	isLeaf := n.Left == nil && n.Right == nil
	if isLeaf {
		fmt.Printf("%s%s%x-%s\033[0m\n", prefix, "\033[32m", n.Data[:], n.DataString)
	} else {
		fmt.Printf("%s%x-%s\n", prefix, n.Data[:], n.DataString)
	}

	if n.Left != nil {
		n.Left.ReatTree(prefix + "│   ")
	}
	if n.Right != nil {
		n.Right.ReatTree(prefix + "    ")
	}
}
