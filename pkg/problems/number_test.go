package problems

import (
	"testing"
)

func TestRational_ToDecimal(t *testing.T) {
	r := Rational{Numerator: 1, Denominator: 8}
	d := r.Decimal()
	if d.String() != "0.125" {
		t.Errorf("unexpected conversion from 1/8 to %s", d.String())
	}

	r = Rational{Numerator: 1, Denominator: 29}
	d = r.Decimal()
	if d.String() != "0.(0344827586206896551724137931)" {
		t.Errorf("unexpected conversion from 1/29 to %s", d.String())
	}
}
