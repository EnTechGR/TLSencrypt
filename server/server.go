package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

func main() {
	// Load the TLS certificate and key
	cert, err := tls.LoadX509KeyPair("../server.crt", "../server.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	// Secure TLS configuration
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12, // Enforce at least TLS 1.2
	}

	// Start a TLS listener
	listener, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}
	defer listener.Close()
	log.Println("TLS Server listening on port 8443...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read client message
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Printf("read error: %v", err)
		return
	}
	log.Printf("Received: %s", buf[:n])

	// Respond to client
	response := "Hello from TLS Server"
	conn.Write([]byte(response))
}
