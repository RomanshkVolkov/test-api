package repository

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"golang.org/x/crypto/blake2b"
)

func ValidateError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func MaskString(s string) string {
	lenght := len(s)
	return strings.Repeat("*", lenght)
}

func TxtToHash(s string) string {
	hash := blake2b.Sum512([]byte(s))
	hexHash := hex.EncodeToString(hash[:])
	return hexHash
}

func TxtToRandomNumbers(s string) string {
	hexHash := TxtToHash(s)
	fmt.Println(hexHash)

	hashBytes, err := hex.DecodeString(hexHash)
	if err != nil {
		panic(err)
	}

	// Crear un lector de bytes a partir del hash
	reader := bytes.NewReader(hashBytes)

	// Usar el hash como semilla para un PRNG
	var seed int64
	err = binary.Read(reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}

	// validar que la semilla sea positiva
	if seed < 0 {
		seed = -seed
	}

	rng := rand.Reader
	fmt.Println("Seed:")
	fmt.Println(seed)
	_, err = rand.Int(rng, big.NewInt(seed))
	if err != nil {
		panic(err)
	}

	randomNumber, err := rand.Int(rng, big.NewInt(1e18))
	if err != nil {
		panic(err)
	}

	return randomNumber.String()
}

func CurrentTime() string {
	return time.Now().UTC().Format("2024-01-01 13:01:01")
}
