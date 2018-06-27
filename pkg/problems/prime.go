package problems

import (
	"math"
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
		next := p.sieveOfEratosthenes()
		<-next
	}
}

func (p *PrimeGenerator) sieveOfEratosthenes() <-chan bool {
	done := make(chan bool)

	go func() {
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
		close(done)
	}()
	return done
}

func (p *PrimeGenerator) IsPrime(n uint64) bool {
	// If we luck out and have already generated all
	// the primes up to n
	if n < p.n {
		return p.isPrime[n]
	}

	limit := uint64(math.Sqrt(float64(n)))
	// value was larger than the current list, but the sqrt of
	// the value is within reach of the list.  Searching
	// if generated primes are factors is usually faster
	// than generating primes up to n
	if limit < p.n {
		idx := 0
		for {
			prime := p.Get(idx)
			if prime > limit {
				break
			}
			if n%prime == 0 {
				return false
			}
			idx++
		}
		return true
	}

	// don't have enough primes generated to tell. generating
	// up to `limit` and launching a recheck
	for {
		prime := p.Get(len(p.primes))
		if prime > limit {
			break
		}
	}
	return p.IsPrime(n)
}
