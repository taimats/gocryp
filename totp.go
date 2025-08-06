package gocrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

const timeStepSecond int64 = 30 //seconds

func TOTP(secretKey []byte, c Clock, digitSize int) string {
	num := hotp(secretKey, counter(c), digitSize)
	digitFmt := fmt.Sprintf("%%0%dd", digitSize)
	return fmt.Sprintf(digitFmt, num)
}

func counter(c Clock) uint64 {
	return uint64(c.Now().Unix() / timeStepSecond)
}

func hotp(secretKey []byte, counter uint64, digitSize int) uint32 {
	hs := hmacSha1(secretKey, counter)
	num := dynamicTruncate(hs)
	return num % uint32(math.Pow10(digitSize))
}

func hmacSha1(secretKey []byte, counter uint64) []byte {
	mac := hmac.New(sha1.New, secretKey)

	c := make([]byte, 8)
	binary.BigEndian.PutUint64(c, counter)

	mac.Write(c)
	return mac.Sum(nil)
}

func dynamicTruncate(value []byte) uint32 {
	offset := value[len(value)-1] & 0xf
	return binary.BigEndian.Uint32(value[offset:offset+4]) & 0x7fffffff
}

type Clock interface {
	Now() time.Time
}

type Clocker struct{}

func NewClocker() *Clocker {
	return &Clocker{}
}

func (c *Clocker) Now() time.Time {
	return time.Now()
}

type PseudoClocker struct{}

func NewPseudoClocker() *PseudoClocker {
	return &PseudoClocker{}
}

func (c *PseudoClocker) Now() time.Time {
	return time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
}
