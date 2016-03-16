package math

import "fmt"

// Fraction represents ratio.
type Fraction struct {
	Numerator, Denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Reduce makes given fraction reduced.
func (f Fraction) Reduce() Fraction {
	gcd := GCDEuclidean(f.Numerator, f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd

	return f
}

// MultiplyByNumber multiplies fraction by given number.
func (f Fraction) MultiplyByNumber(m int) Fraction {
	f.Numerator *= m

	return f
}

// MultiplyByFraction multiplies fraction by given fraction.
func (f Fraction) MultiplyByFraction(m Fraction) Fraction {
	f.Numerator *= m.Numerator
	f.Denominator *= m.Denominator

	return f
}
