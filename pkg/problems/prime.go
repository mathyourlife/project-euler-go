package problems

import (
// "log"
)

var (
	pg *PrimeGenerator
)

func init() {
	pg = NewPrimeGenerator()
}

func GetPrime(n int) uint64 {
	return pg.Get(n)
}

type PrimeGenerator struct {
	primes []uint64
	sieve  []uint64
	n      uint64
}

func NewPrimeGenerator() *PrimeGenerator {
	p := &PrimeGenerator{
		primes: make([]uint64, 0, 100),
		sieve:  make([]uint64, 0, 100),
	}
	p.primes = append(p.primes, 2, 3)
	p.sieve = append(p.sieve, 2, 3)
	p.n = 3
	return p
}

func (p *PrimeGenerator) Get(n int) uint64 {
	for {
		if n < len(p.primes) {
			return p.primes[n]
		}
		p.sieveOfEratosthenes()
	}
}

func (p *PrimeGenerator) sieveOfEratosthenes() {
	var isPrime bool

	for {
		// log.Println("----------------")
		// log.Println("n", p.n)
		// log.Println("primes", p.primes)
		// log.Println("sieve", p.sieve)
		p.n += 2
		isPrime = true
	SieveLoop:
		for i := 0; i < len(p.sieve); i++ {
			for {
				if p.sieve[i] > p.n {
					break
				} else if p.sieve[i] == p.n {
					isPrime = false
					break SieveLoop
				}
				p.sieve[i] += p.primes[i]
			}
		}
		if isPrime {
			break
		}
	}
	p.primes = append(p.primes, p.n)
	p.sieve = append(p.sieve, p.n)
}
