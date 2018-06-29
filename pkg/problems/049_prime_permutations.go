package problems

import (
	"fmt"
)

type PrimePermutations struct{}

func (p *PrimePermutations) ID() int {
	return 49
}

func (p *PrimePermutations) Text() string {
	return `The arithmetic sequence, 1487, 4817, 8147, in which each
of the terms increases by 3330, is unusual in two ways:

(i) each of the three terms are prime, and,
(ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes,
exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three
terms in this sequence?
`
}

func (p *PrimePermutations) Solve() (string, error) {

	// helper function to create a number from the corresponding
	// digits int slice.
	// ex:
	//   digits = []int{3,2,4}
	//   returns 324
	sliceToNumber := func(digits []int) int {
		n := 0
		for _, d := range digits {
			n = n*10 + d
		}
		return n
	}

	genSolution := func(ns []int) string {
		s := ""
		for i := 0; i < len(ns)-1; i++ {
			for j := i + 1; j < len(ns); j++ {
				if ns[j]-ns[i] == 3330 {
					if len(s) == 0 {
						s = fmt.Sprintf("%d%d", ns[i], ns[j])
					} else {
						s += fmt.Sprintf("%d", ns[j])
					}
				}
			}
		}
		return s
	}

	primes := make(chan int)

	go func() {
		n := 0
		defer close(primes)
		for {
			n++
			v := int(GetPrime(n))
			if v < 1000 {
				continue
			}
			if v >= 10000 {
				return
			}
			primes <- v
		}
	}()

	for v := range primes {
		ns := []int{}

		digits := []int{v / 1000 % 10, v / 100 % 10, v / 10 % 10, v % 10}
		for {
			num := sliceToNumber(digits)
			if IsPrime(uint64(num)) {
				ns = append(ns, num)
			}
			if !LexPerm(digits) {
				break
			}
		}
		if len(ns) < 3 {
			continue
		}
		diffs := map[int]bool{}
		for i := 0; i < len(ns)-1; i++ {
			for j := i + 1; j < len(ns); j++ {
				if ns[j]-ns[i] != 3330 {
					continue
				}
				if ns[j] == 1487 || ns[j] == 4817 || ns[j] == 8147 {
					continue
				}
				if diffs[ns[j]-ns[i]] {
					return genSolution(ns), nil
				}
				diffs[ns[j]-ns[i]] = true
			}
		}
	}

	return fmt.Sprintf("%d", 0), nil
}
