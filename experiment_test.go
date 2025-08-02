package gocrypt_test

import "testing"

func BenchmarkEfficientPow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		efficientPow(3, 10)
	}
}

func BenchmarkNonEffPow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		nonEffPow(3, 10)
	}
}

func efficientPow(v int, p int) int {
	size := bitSize(p)
	num := 1
	mulNum := v
	for range size {
		if (p & 1) == 1 {
			num *= mulNum
		}
		mulNum *= mulNum
		p >>= 1
	}
	return num
}

func nonEffPow(v int, p int) int {
	num := 1
	for range p {
		num *= v
	}
	return num
}

func bitSize(p int) int {
	n := 64
	for i := n - 1; i >= 0; i-- {
		if (p >> i & 1) == 1 {
			return i + 1
		}
	}
	return 0
}

func lowestBit(v int, shitNum int) int {
	return v >> shitNum & 1
}

func TestEfficientPow(t *testing.T) {
	want := 59049
	got := efficientPow(3, 10)
	if got != want {
		t.Errorf("EfficientPow: Not equal:\n(got=%d, want=%d)\n", got, want)
	}
}

func TestNoEffPow(t *testing.T) {
	want := 59049
	got := nonEffPow(3, 10)
	if got != want {
		t.Errorf("NoEffPow: Not equal:\n(got=%d, want=%d)\n", got, want)
	}
}
