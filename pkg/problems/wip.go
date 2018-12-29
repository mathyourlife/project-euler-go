// This is a dumping ground for stuff that is not in use, but
// may be useful in the future
package problems

// nCr pick the next largest combination
// of indicies to retrieve primes.
func next(state []int) {
	for i := 1; i < len(state); i++ {
		v := state[i]
		if v-state[i-1] != 1 {
			state[i-1]++
			for j := 0; j < i-1; j++ {
				state[j] = j
			}
			return
		}
	}
	state[len(state)-1]++
	for i := 0; i < len(state)-1; i++ {
		state[i] = i
	}
}

func combinations(n, r int) [][]int {
	cs := [][]int{}
	if r > n {
		return cs
	}
	idx := make([]int, 0, r)
	for i := 0; i < r; i++ {
		idx = append(idx, i)
	}

	c := []int{}
	for _, i := range idx {
		c = append(c, i)
	}
	cs = append(cs, c)

	var i int
	for {
		done := true
		for i = r - 1; i >= 0; i-- {
			if idx[i] != i+n-r {
				done = false
				break
			}
		}
		if done {
			return cs
		}
		idx[i]++
		for j := i + 1; j < r; j++ {
			idx[j] = idx[j-1] + 1
		}
		c := []int{}
		for _, i := range idx {
			c = append(c, i)
		}
		cs = append(cs, c)
	}
	return cs
}
