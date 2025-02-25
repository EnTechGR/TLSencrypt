package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Load CA certificate
	caCert, err := os.ReadFile("../server.crt") // Use a CA certificate, not a self-signed cert in production
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}

	// Create a certificate pool and add the CA certificate
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("failed to add CA certificate to pool")
	}

	// Secure TLS configuration
	config := &tls.Config{
		RootCAs:    caCertPool, // Trust the specified CA
		MinVersion: tls.VersionTLS12,
	}

	// Connect to the TLS server
	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to TLS Server")

	// Send a message
	message := "Hello TLS Server"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("write error: %v", err)
	}

	// Read server response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalf("read error: %v", err)
	}
	fmt.Printf("Server response: %s\n", buf[:n])
}
