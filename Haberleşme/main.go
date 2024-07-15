package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
	mu     sync.Mutex
}

type Node struct {
	Address string
}

type Network struct {
	Nodes []*Node
}

func (b *Block) CalculateHash() string {
	// Burada gerçek bir hash hesaplama fonksiyonu kullanılmalıdır.
	// Örnek olarak basit bir hash döndürüyoruz.
	return fmt.Sprintf("%d%d%s", b.Index, b.Timestamp, b.PrevHash)
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.Hash = newBlock.CalculateHash()
	bc.mu.Lock()
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.mu.Unlock()
}

func (bc *Blockchain) CreateGenesisBlock() *Block {
	return &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		PrevHash:  "0",
		Hash:      "0",
	}
}

func (n *Node) Start(bc *Blockchain, network *Network) {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		fmt.Println("Error starting node:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Node started at", n.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, bc, network)
	}
}

func handleConnection(conn net.Conn, bc *Blockchain, network *Network) {
	defer conn.Close()

	var msg map[string]interface{}
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	switch msg["type"] {
	case "transaction":
		var tx Transaction
		txJSON, _ := json.Marshal(msg["data"])
		json.Unmarshal(txJSON, &tx)
		bc.mu.Lock()
		bc.Blocks[len(bc.Blocks)-1].Transactions = append(bc.Blocks[len(bc.Blocks)-1].Transactions, tx)
		bc.mu.Unlock()
	case "block":
		var block Block
		blockJSON, _ := json.Marshal(msg["data"])
		json.Unmarshal(blockJSON, &block)
		bc.mu.Lock()
		bc.Blocks = append(bc.Blocks, &block)
		bc.mu.Unlock()
	}
}

func (n *Node) Broadcast(msg map[string]interface{}, network *Network) {
	for _, node := range network.Nodes {
		if node.Address != n.Address {
			conn, err := net.Dial("tcp", node.Address)
			if err != nil {
				fmt.Println("Error connecting to node:", err)
				continue
			}
			encoder := json.NewEncoder(conn)
			encoder.Encode(msg)
			conn.Close()
		}
	}
}

func main() {
	bc := &Blockchain{}
	bc.Blocks = append(bc.Blocks, bc.CreateGenesisBlock())

	network := &Network{}
	network.Nodes = append(network.Nodes, &Node{Address: "localhost:8000"})
	network.Nodes = append(network.Nodes, &Node{Address: "localhost:8001"})

	for _, node := range network.Nodes {
		go node.Start(bc, network)
	}

	// Örnek işlem ve blok oluşturma
	time.Sleep(2 * time.Second)
	tx := Transaction{Sender: "Alice", Receiver: "Bob", Amount: 1.5}
	msg := map[string]interface{}{"type": "transaction", "data": tx}
	network.Nodes[0].Broadcast(msg, network)

	time.Sleep(2 * time.Second)
	bc.AddBlock([]Transaction{tx})
	blockJSON, _ := json.Marshal(bc.Blocks[len(bc.Blocks)-1])
	msg = map[string]interface{}{"type": "block", "data": json.RawMessage(blockJSON)}
	network.Nodes[0].Broadcast(msg, network)

	time.Sleep(2 * time.Second)
	fmt.Println("Blockchain:", bc.Blocks)
	printBlockchain(bc)
}

func printBlockchain(bc *Blockchain) {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Index: %d\n", block.Index)
		fmt.Printf("  Timestamp: %d\n", block.Timestamp)
		fmt.Printf("  PrevHash: %s\n", block.PrevHash)
		fmt.Printf("  Hash: %s\n", block.Hash)
		fmt.Printf("  Transactions:\n")
		for j, tx := range block.Transactions {
			fmt.Printf("    Transaction %d:\n", j)
			fmt.Printf("      Sender: %s\n", tx.Sender)
			fmt.Printf("      Receiver: %s\n", tx.Receiver)
			fmt.Printf("      Amount: %.2f\n", tx.Amount)
		}
	}
}
