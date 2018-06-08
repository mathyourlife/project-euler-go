package problems

import (
	"fmt"
)

type CoinSums struct {
	count int
}

func (p *CoinSums) ID() int {
	return 31
}

func (p *CoinSums) Text() string {
	return `In England the currency is made up of pound, £,
and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
`
}

// given a map of [string]int return the keys in
// increasing order by value
func (c *CoinSums) sortMapByValue(m map[string]int) []string {
	var min int

	// Pick a random value from the map as the initial minimum
	for _, v := range m {
		min = v
		break
	}
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	min--

	keys := make([]string, 0, len(m))
	for i := 0; i < len(m); i++ {
		key := ""
		value := min
		for k, v := range m {
			if v > min && (value == min || v < value) {
				key = k
				value = v
			}
		}
		keys = append(keys, key)
		min = value
	}
	return keys
}

func (p *CoinSums) check(a, b, c, d, e, f, g, h int) int {
	total := a*1 + b*2 + c*5 + d*10 + e*20 + f*50 + g*100 + h*200
	if total == 200 {
		p.count++
	}
	return total
}

func (p *CoinSums) Solve() (string, error) {

	goal := 200
	coin := map[string]int{
		"1p":  1,
		"2p":  2,
		"5p":  5,
		"10p": 10,
		"20p": 20,
		"50p": 50,
		"£1":  100,
		"£2":  200,
	}

	// order names from least to greatest
	names := p.sortMapByValue(coin)

	purse := make([]int, 0, len(coin))
	for _ = range coin {
		purse = append(purse, 0)
	}

	for a := 0; a <= goal/coin[names[0]]; a++ {
		if p.check(a, 0, 0, 0, 0, 0, 0, 0) >= goal {
			break
		}
		for b := 0; b <= goal/coin[names[1]]; b++ {
			if p.check(a, b, 0, 0, 0, 0, 0, 0) >= goal {
				break
			}
			for c := 0; c <= goal/coin[names[2]]; c++ {
				if p.check(a, b, c, 0, 0, 0, 0, 0) >= goal {
					break
				}
				for d := 0; d <= goal/coin[names[3]]; d++ {
					if p.check(a, b, c, d, 0, 0, 0, 0) >= goal {
						break
					}
					for e := 0; e <= goal/coin[names[4]]; e++ {
						if p.check(a, b, c, d, e, 0, 0, 0) >= goal {
							break
						}
						for f := 0; f <= goal/coin[names[5]]; f++ {
							if p.check(a, b, c, d, e, f, 0, 0) >= goal {
								break
							}
							for g := 0; g <= goal/coin[names[6]]; g++ {
								if p.check(a, b, c, d, e, f, g, 0) >= goal {
									break
								}
								for h := 0; h <= goal/coin[names[7]]; h++ {
									if p.check(a, b, c, d, e, f, g, h) >= goal {
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return fmt.Sprintf("%d", p.count), nil
}
