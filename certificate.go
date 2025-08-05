package gocrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	leastBitSize   = 1024
	defaultBitSize = 2048
)

const defaultDir = "secret"

var (
	ErrInvalidBitSize = errors.New("bit size must be more than 1024")
)

func PemKeyPairs(args []string) error {
	bitsize, err := handleArgs(args)
	if err != nil {
		return err
	}
	privKey, err := pemPrivateKey(bitsize)
	if err != nil {
		return err
	}
	err = pemPublicKey(privKey)
	if err != nil {
		return err
	}
	return nil
}

func handleArgs(args []string) (bitSize int, err error) {
	if len(args) != 2 {
		return defaultBitSize, nil
	}
	bitSize, err = strconv.Atoi(args[1])
	if err != nil {
		return 0, err
	}
	if bitSize < leastBitSize {
		return 0, ErrInvalidBitSize
	}
	return bitSize, nil
}

func pemPrivateKey(bits int) (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(fmt.Sprintf("%s/private.pem", defaultDir))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	privPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	if err := pem.Encode(f, privPem); err != nil {
		return nil, err
	}
	fmt.Println("Success: a secret key in pem is generated.")

	return key, nil
}

func pemPublicKey(key *rsa.PrivateKey) error {
	f, err := os.Create(fmt.Sprintf("%s/public.pem", defaultDir))
	if err != nil {
		return err
	}
	defer f.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return err
	}
	pubPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	if err := pem.Encode(f, pubPem); err != nil {
		return err
	}
	fmt.Println("Success: a public key in pem is generated.")

	return nil
}
