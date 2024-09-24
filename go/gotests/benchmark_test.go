package unit

import (
	"testing"
	"unit_test/funcs"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcs.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcs.MakeSliceWithAlloc()
	}
}
