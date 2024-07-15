package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type Node struct {
	Address string
	Peers   map[string]bool
	mu      sync.Mutex
}

func NewNode(address string) *Node {
	return &Node{
		Address: address,
		Peers:   make(map[string]bool),
	}
}

func (n *Node) Start() {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		fmt.Println("Error starting node:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Node started at", n.Address)

	go n.handleInput()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("New connection from", remoteAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("Received message:", msg)
		n.broadcast(msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}

	fmt.Println("Connection closed from", remoteAddr)
}

func (n *Node) handleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		n.broadcast(msg)
	}
}

func (n *Node) broadcast(msg string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for peer := range n.Peers {
		go n.sendMessage(peer, msg)
	}
}

func (n *Node) sendMessage(peer, msg string) {
	conn, err := net.Dial("tcp", peer)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}
	defer conn.Close()

	_, err = fmt.Fprintln(conn, msg)
	if err != nil {
		fmt.Println("Error sending message to peer:", err)
	}
}

func (n *Node) ConnectToPeer(peer string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if _, exists := n.Peers[peer]; !exists {
		n.Peers[peer] = true
		go n.sendMessage(peer, "Hello from "+n.Address)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <node_address>")
		return
	}

	nodeAddress := os.Args[1]
	node := NewNode(nodeAddress)

	go node.Start()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter peer address to connect: ")
		peer, _ := reader.ReadString('\n')
		peer = strings.TrimSpace(peer)
		node.ConnectToPeer(peer)
	}
}

/*
Kodu çalıştırmak için, terminalde aşağıdaki komutu kullanın:


go run main.go <node_address>



Copy code
go run main.go localhost:8000
Ardından, diğer nodları bağlamak için eş adreslerini girin. Örneğin:


Enter peer address to connect: localhost:8001
Bu şekilde, nodlar birbirleriyle iletişim kuracak ve mesajları paylaşacaktır.
*/

/*
Peer-to-Peer (P2P) ağının ana olayı, merkezi bir sunucuya ihtiyaç duymadan, ağdaki tüm düğümlerin (eşlerin) birbirleriyle doğrudan iletişim kurmaları ve kaynakları paylaşmalarıdır. Bu, ağın dağıtık ve merkeziyetsiz bir yapıda çalışmasını sağlar.

P2P ağının ana özellikleri şunlardır:

Dağıtık Yapı: P2P ağında, tüm düğümler eşit statüdedir ve her düğüm hem istemci hem de sunucu görevi görür. Bu, ağın merkezi bir sunucuya ihtiyaç duymadan çalışmasını sağlar.

Kaynak Paylaşımı: Düğümler, dosyalar, bant genişliği, işlem gücü ve diğer kaynakları birbirleriyle paylaşabilirler. Bu, ağın ölçeklenebilirliğini ve verimliliğini artırır.

Esneklik ve Dayanıklılık: P2P ağı, merkezi bir sunucuya bağlı olmadığı için, tek bir hata noktası olmadığından daha dayanıklıdır. Ağdaki bir düğümün arızalanması, ağın geri kalanını etkilemez.

Gizlilik ve Güvenlik: P2P ağında, düğümler doğrudan iletişim kurar ve merkezi bir otorite tarafından izlenmezler. Bu, kullanıcıların gizliliğini ve güvenliğini artırabilir.

Ölçeklenebilirlik: P2P ağı, yeni düğümlerin eklenmesiyle doğal olarak ölçeklenebilir. Ağdaki düğüm sayısı arttıkça, ağın kapasitesi ve performansı da artar.

Maliyet Tasarrufu: P2P ağı, merkezi bir sunucuya ihtiyaç duymadığı için, altyapı maliyetlerini azaltır ve düşük maliyetli bir şekilde çalışabilir.

P2P ağları, dosya paylaşımı, işlem gücü paylaşımı, canlı yayın, blockchain ve diğer birçok uygulamada kullanılmaktadır. Bu ağlar, merkeziyetsizlik, esneklik ve dayanıklılık gibi avantajları nedeniyle giderek daha popüler hale gelmektedir.
*/
