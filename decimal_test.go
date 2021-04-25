package decimal

import "testing"
import dec "github.com/shopspring/decimal"

func TestDecimal(t *testing.T) {
	d1, _ := dec.NewFromString("0.23")
	d2 := dec.NewFromInt(123)
	d3 := dec.NewFromFloat(309.23)

	a := d1.Add(d2).Mul(d3).Sub(d1).Div(d2)

	f1, _ := NewFromString("0.23")
	f2 := NewFromInt(123)
	f3 := NewFromFloat(309.23)

	b := f1.Add(f2).Mul(f3).Sub(f1).Div(f2)

	t.Log(a, b)
}

func BenchmarkDecimal1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1, _ := dec.NewFromString("0.23")
		d2 := dec.NewFromInt(123)
		d3 := dec.NewFromFloat(309.23)

		d1.Add(d2).Mul(d3).Sub(d1).Div(d2)
	}
}

func BenchmarkDecimal2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1, _ := NewFromString("0.23")
		f2 := NewFromInt(123)
		f3 := NewFromFloat(309.23)

		f1.Add(f2).Mul(f3).Sub(f1).Div(f2)
	}
}