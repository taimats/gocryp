package main

import (
	"fmt"
	"os"

	gocrypt "github.com/taimats/gocryp"
)

func main() {
	err := gocrypt.PemKeyPairs(os.Args)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
