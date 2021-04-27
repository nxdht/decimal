package decimal

import "testing"

func BenchmarkTestDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFromString("0.2345")
	}
}
