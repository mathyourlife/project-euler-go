package problems

import (
	"fmt"
)

type AmicableNumbers struct{}

func (p *AmicableNumbers) ID() int {
	return 21
}

func (p *AmicableNumbers) Text() string {
	return `Let d(n) be defined as the sum of proper divisors of n
(numbers less than n which divide evenly into n).

If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable
pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are
1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;

therefore d(220) = 284. The proper divisors of 284 are
1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
`
}

func (p *AmicableNumbers) Solve() (string, error) {

	sum := uint64(0)
	for n := uint64(1); n < uint64(10000); n++ {
		a := p.amicable(n)
		if n != a && n == p.amicable(a) {
			sum += n
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func (p *AmicableNumbers) amicable(n uint64) uint64 {
	pf := primeFactors(n)

	divisors := []uint64{1}
	for factor, exponent := range pf {
		multiple := uint64(1)
		newDivisors := []uint64{}
		for e := 1; e <= exponent; e++ {
			multiple *= factor
			for _, divisor := range divisors {
				newDivisors = append(newDivisors, divisor*multiple)
			}
		}
		divisors = append(divisors, newDivisors...)
	}

	amicable := uint64(0)
	for _, v := range divisors[:len(divisors)-1] {
		amicable += v
	}
	return amicable
}
