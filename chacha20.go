package gocrypt

import (
	"crypto/cipher"
	"encoding/binary"
)

type Cipher struct {
	constant [4]uint32
	key      [8]uint32
	counter  uint32
	nonce    [3]uint32
}

var _ cipher.Stream = (*Cipher)(nil)

var constants = [4]uint32{0x61707865, 0x3320646e, 0x79622d32, 0x6b206574}

func NewCipher(key [32]byte, count uint32, nonce [12]byte) *Cipher {
	c := new(Cipher)
	c.constant = constants
	pos := 4
	for i := range 8 {
		if pos <= len(key) {
			break
		}
		c.key[i] = binary.LittleEndian.Uint32(key[pos-4 : pos])
		pos += 4
	}
	c.counter = count
	c.nonce = [3]uint32{
		binary.LittleEndian.Uint32(nonce[0:4]),
		binary.LittleEndian.Uint32(nonce[4:8]),
		binary.LittleEndian.Uint32(nonce[8:12]),
	}
	return c
}

func (c *Cipher) toState() [16]uint32 {
	return [16]uint32{
		c.constant[0], c.constant[1], c.constant[2], c.constant[3],
		c.key[0], c.key[1], c.key[2], c.key[3],
		c.key[4], c.key[5], c.key[6], c.key[7],
		c.counter, c.nonce[0], c.nonce[1], c.nonce[2],
	}
}

func (c *Cipher) XORKeyStream(dst, src []byte) {

}

func rotateLeft32(x uint32, k int) uint32 {
	const n = 32
	if k >= 0 {
		s := uint(k)
		return x<<s | x>>(n-s)
	}
	s := uint(-k)
	return x>>s | x<<(n-s)
}
