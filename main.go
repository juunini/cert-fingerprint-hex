package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		certPath string
		output   string
	)
	flag.StringVar(&certPath, "cert", "cert.pem", "Path to the certificate file")
	flag.StringVar(&output, "output", "fingerprint.hex", "Path to the output file")
	flag.Parse()

	pemData, err := os.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		fmt.Println("PEM decoding failed")
		os.Exit(1)
	}

	derData := block.Bytes

	hasher := sha256.New()
	hasher.Write(derData)
	hash := hasher.Sum(nil)

	fingerprint, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer fingerprint.Close()

	fingerprint.WriteString(hex.EncodeToString(hash))
}
