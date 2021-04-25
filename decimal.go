package decimal

import "github.com/robaho/fixed"

var Zero = Decimal{fixed.ZERO}

type Decimal struct {
	fixed.Fixed
}

func NewFromInt(value int64) Decimal {
	return Decimal{fixed.NewI(value, 0)}
}

func NewFromString(value string) (Decimal, error) {
	f, e := fixed.NewSErr(value)
	if e != nil {
		return Decimal{}, e
	}
	return Decimal{f}, nil
}

func NewFromFloat(value float64) Decimal {
	return Decimal{fixed.NewF(value)}
}

func (d Decimal) Abs() Decimal {
	return Decimal{d.Fixed.Abs()}
}

func (d Decimal) Add(d2 Decimal) Decimal {
	return Decimal{d.Fixed.Add(d2.Fixed)}
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	return Decimal{d.Fixed.Sub(d2.Fixed)}
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal{d.Fixed.Mul(d2.Fixed)}
}

func (d Decimal) Div(d2 Decimal) Decimal {
	return Decimal{d.Fixed.Div(d2.Fixed)}
}

func (d Decimal) Equal(d2 Decimal) bool {
	return d.Fixed.Equal(d2.Fixed)
}

func (d Decimal) Cmp(d2 Decimal) int {
	return d.Fixed.Cmp(d2.Fixed)
}

func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Cmp(d2) == 1
}

func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == 1 || cmp == 0
}

func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Cmp(d2) == -1
}

func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == -1 || cmp == 0
}

func (d Decimal) IsPositive() bool {
	return d.Fixed.Sign() == 1
}

func (d Decimal) IsNegative() bool {
	return d.Fixed.Sign() == -1
}

func (d Decimal) IsZero() bool {
	return d.Fixed.IsZero()
}

func (d Decimal) Floor() Decimal {
	return Decimal{d.Fixed.Floor()}
}

func (d Decimal) Int() Decimal {
	return Decimal{fixed.NewI(d.Fixed.Int(), 0)}
}

func Min(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}

func Max(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

func Sum(first Decimal, rest ...Decimal) Decimal {
	total := first
	for _, item := range rest {
		total = total.Add(item)
	}

	return total
}

func Avg(first Decimal, rest ...Decimal) Decimal {
	count := NewFromInt(int64(len(rest)+1))
	sum := Sum(first, rest...)
	return sum.Div(count)
}


func (d Decimal) String() string {
	return d.Fixed.String()
}

func (d Decimal) StringFixed(places int32) string {
	return d.Fixed.StringN(int(places))
}