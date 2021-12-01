package util

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	fmt.Println(RandStr(32))
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandStr(32)
	}
}