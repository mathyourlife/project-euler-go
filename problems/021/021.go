/*
Let d(n) be defined as the sum of proper divisors of n (numbers less
than n which divide evenly into n).

If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable
pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are
1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;

therefore d(220) = 284. The proper divisors of 284 are
1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"
	"math"
)

type PrimeFactorization struct {
	cache map[uint64]map[uint64]uint64
}

func NewPrimeFactorization() *PrimeFactorization {

	return &PrimeFactorization{
		cache: make(map[uint64]map[uint64]uint64),
	}
}

func (pf *PrimeFactorization) Of(n uint64) map[uint64]uint64 {

	if len(pf.cache[n]) > 0 {
		return pf.cache[n]
	}

	var t bool

	orig := n
	factor := uint64(2)
	f := [][]uint64{}

	use_cache := false
	for {
		if len(pf.cache[n]) > 0 {
			use_cache = true
			fs := map[uint64]uint64{}
			for k, v := range pf.cache[n] {
				fs[k] = v
			}
			for i := len(f) - 1; i >= 0; i-- {
				fs[f[i][1]]++
				pf.cache[f[i][0]] = map[uint64]uint64{}
				for k, v := range fs {
					pf.cache[f[i][0]][k] += v
				}
			}
			break
		}
		t, n = pf.is_a_factor(n, factor)
		if t {
			f = append(f, []uint64{n * factor, factor})
			if n <= 1 {
				break
			}
		} else {
			factor++
		}
		if n <= 1 {
			break
		}
	}
	if !use_cache {
		pf.cache_it(f)
	}
	return pf.cache[orig]
}

func (pf *PrimeFactorization) is_a_factor(composite uint64, factor uint64) (bool, uint64) {
	if composite < 2 {
		return false, composite
	}
	if composite%factor == 0 {
		return true, composite / factor
	}
	return false, composite
}

func (pf *PrimeFactorization) cache_it(f [][]uint64) {
	fs := map[uint64]uint64{}
	for i := len(f) - 1; i >= 0; i-- {
		fs[f[i][1]]++
		pf.cache[f[i][0]] = map[uint64]uint64{}
		for k, v := range fs {
			pf.cache[f[i][0]][k] = v
		}
	}
}

func divisors(n uint64, pf *PrimeFactorization) []uint64 {

	ds := map[uint64]bool{
		1: true,
	}
	for k, v := range pf.Of(n) {
		new_divisors := []uint64{}
		for divisor, _ := range ds {
			for i := uint64(0); i < v; i++ {
				multiple := k * uint64(math.Pow(float64(k), float64(i)))
				new_divisors = append(new_divisors, divisor*multiple)
			}
		}
		for _, divisor := range new_divisors {
			ds[divisor] = true
		}
	}

	list := []uint64{}
	for n, _ := range ds {
		list = append(list, n)
	}
	return list
}

func main() {
	N := uint64(10000)

	pf := NewPrimeFactorization()

	// Load prime factorizations for numbers up to N
	for i := uint64(2); i < N; i++ {
		pf.Of(i)
	}

	asums := map[uint64]uint64{}
	for i := uint64(2); i < N; i++ {
		ds := divisors(i, pf)
		sum := uint64(0)
		for _, d := range ds {
			if d != i {
				sum += d
			}
		}
		asums[i] = sum
	}
	sum := uint64(0)
	for a, b := range asums {
		if a == b {
			continue
		}
		if a > b {
			// skip double reporting
			continue
		}
		if a == asums[b] && asums[b] == a {
			sum += a
			sum += b
		}
	}
	fmt.Println(sum)
}
