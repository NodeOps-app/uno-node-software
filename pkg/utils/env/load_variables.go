package env

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
	"strconv"
)

func MustLoadString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal(key, " environment variable is required")
	}
	return val
}

func MustLoadUint64(key string) uint64 {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal(key, " environment variable is required")
	}

	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %s as uint64: %v", key, err)
	}

	return num
}

func MustLoadPrivateKey(key string) *ecdsa.PrivateKey {
	val := MustLoadString(key)
	if pk, err := crypto.HexToECDSA(val); err != nil {
		log.Fatal(key, " is invalid")
	} else {
		return pk
	}
	return nil
}
