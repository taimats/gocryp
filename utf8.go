package gocrypt

import (
	"bytes"
)

type encodedSize uint8

const (
	oneByte encodedSize = iota + 1
	twoByte
	threeByte
)

// エンコードサイズに応じたエンコードを行う関数
type varEncodeFunc func(r rune) []byte

type UTF8Encoder struct {
	text  string
	encFn map[encodedSize]varEncodeFunc
}

func NewUTF8Encoder(text string) *UTF8Encoder {
	m := make(map[encodedSize]varEncodeFunc)
	m[oneByte] = encodeOneByte
	m[twoByte] = encodeTwoByte
	m[threeByte] = encodeThreeByte

	return &UTF8Encoder{
		text:  text,
		encFn: m,
	}
}

func (enc *UTF8Encoder) Encode() []byte {
	var buf bytes.Buffer
	for _, r := range enc.text {
		size := enc.checkEncodedSize(r)
		if size == 0 {
			return nil
		}
		buf.Write(enc.encFn[size](r))
	}
	return buf.Bytes()
}

// unicode(16進数表記)の値に応じて、エンコード後のサイズを決定。
// 1バイト: U+0000からU+007F（ASCII）
// 2バイト: U+0080からU+07FF
// 3バイト: U+0800からU+FFFF
// 4バイト: U+10000からU+10FFFF
// 上記に当てはまらない文字の場合は0を返す（つまり、エンコードできない）。
func (enc *UTF8Encoder) checkEncodedSize(r rune) encodedSize {
	if 0x00 <= r || r <= 0x7f {
		return oneByte
	}
	if 0x80 <= r || r <= 0x7ff {
		return twoByte
	}
	if 0x800 <= r || r <= 0xffff {
		return threeByte
	}
	return 0
}

func encodeOneByte(r rune) []byte {
	b := make([]byte, oneByte)
	val := r & 0x7f
	return append(b, byte(val))
}

func encodeTwoByte(r rune) []byte {
	b := make([]byte, twoByte)
	val := r & 0xbf
	val |= (r >> 6 & 0xdf) << 8
	return append(b, byte(val))
}

func encodeThreeByte(r rune) []byte {
	b := make([]byte, threeByte)
	val := r & 0xbf               //下位1Byteを 10xxxxxx のフォーマットに
	val |= (r >> 6 & 0xbf) << 8   //下位2Byteを 10xxxxxx 10xxxxxx のフォーマットに
	val |= (r >> 12 & 0xef) << 16 //最終的に 1110xxxx 10xxxxxx 10xxxxxx が完成
	return append(b, byte(val))
}
