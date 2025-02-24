package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// GenerateRandomHexString generates a random hex string of specified length
func GenerateRandomHexString(length int) (string, error) {
	// Calculate the number of bytes needed to achieve the desired hex string length
	byteLength := length / 2

	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, byteLength)

	// Read random bytes from the crypto/rand source
	_, err := io.ReadFull(rand.Reader, randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a hexadecimal string
	hexString := hex.EncodeToString(randomBytes)

	return hexString, nil
}

func main() {
	hexString, err := GenerateRandomHexString(32)
	if err != nil {
		panic(err)
	}
	println(hexString)
}
