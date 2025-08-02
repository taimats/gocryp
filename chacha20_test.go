package gocrypt_test

import (
	"testing"

	gocrypt "github.com/taimats/gocryp"
)

func TestRotateLeft32(t *testing.T) {
	tests := []struct {
		x    uint32
		k    int
		want uint32
	}{
		{15, 2, 60},
		{15, 0, 15},
		{15, -2, 3221225475},
	}
	for _, tt := range tests {
		got := gocrypt.RotateLeft32(tt.x, tt.k)
		if got != tt.want {
			t.Errorf("RotateLeft32: Not Equal: \ngot  = %032b\nwant = %032b\n", got, tt.want)
		}
	}
}
