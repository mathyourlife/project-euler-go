package problems

import ()

var (
	pg *PrimeGenerator
)

func init() {
	pg = NewPrimeGenerator()
}

func GetPrime(n int) uint64 {
	return pg.Get(n)
}

func IsPrime(n uint64) bool {
	return pg.IsPrime(n)
}

type PrimeGenerator struct {
	primes  []uint64
	isPrime map[uint64]bool
	sieve   []uint64
	n       uint64
}

func NewPrimeGenerator() *PrimeGenerator {
	p := &PrimeGenerator{
		primes:  make([]uint64, 0, 100),
		sieve:   make([]uint64, 0, 100),
		isPrime: map[uint64]bool{},
	}
	p.primes = append(p.primes, 2, 3)
	p.sieve = append(p.sieve, 2, 3)
	p.isPrime[2] = true
	p.isPrime[3] = true
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
	p.isPrime[p.n] = true
}

func (p *PrimeGenerator) IsPrime(n uint64) bool {
	for {
		if p.primes[len(p.primes)-1] >= n {
			break
		}
		p.Get(len(p.primes))
	}
	return p.isPrime[n]
}
