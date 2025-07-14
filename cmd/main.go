package main

import (
	"fmt"

	gocrypt "github.com/taimats/gocryp"
)

func main() {
	text := "い"
	utf8 := gocrypt.NewUTF8Encoder(text)
	b := utf8.Encode()
	fmt.Printf("元のい:%b\n", 'い')
	fmt.Printf("元のい:%x\n", 'い')
	fmt.Printf("%b\n", b)
	fmt.Printf("%dバイト\n", len(b))
}
