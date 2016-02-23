package math

import "fmt"

type Fraction struct {
	Numerator, Denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

func (f Fraction) Reduce() Fraction {
	gcd := GCDEuclidean(f.Numerator, f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd

	return f;
}

func (f Fraction) MultiplyByNumber(m int) Fraction {
	f.Numerator *= m

	return f
}
