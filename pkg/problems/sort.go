package problems

// LexPerm Modify the provided slice in place, and return false if the slice
// is already in descending order (no more permutations exist).
func LexPerm(a []int) bool {
	k := lexPermGetK(a)
	if k < 0 {
		return false
	}
	l := lexPermGetL(a, k)
	// Step 3 swap k and l
	a[k], a[l] = a[l], a[k]
	// Step 4 revers items for elements > k+1
	for i, j := k+1, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return true
}

// lexPermGetK Find the largest index k such that a[k] < a[k + 1].
// If no such index exists, the permutation is the last permutation.
//
// Return -1 if items are in descending order.
func lexPermGetK(a []int) int {
	k := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			k = i
		}
	}
	return k
}

// lexPermGetL Find the largest index l greater than k such that a[k] < a[l].
func lexPermGetL(a []int, k int) int {
	var l int
	for i := k + 1; i < len(a); i++ {
		if a[k] < a[i] {
			l = i
		}
	}
	return l
}
