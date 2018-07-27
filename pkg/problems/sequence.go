package problems

func PolynomialSequence(p, n int) uint64 {
	if p == 3 {
		return TriangleSequence(n)
	} else if p == 4 {
		return SquareSequence(n)
	} else if p == 5 {
		return PentagonalSequence(n)
	} else if p == 6 {
		return HexagonalSequence(n)
	} else if p == 7 {
		return HeptagonalSequence(n)
	} else if p == 8 {
		return OctagonalSequence(n)
	}
	panic("invalid sequence type")
}

func TriangleSequence(n int) uint64 {
	m := uint64(n)
	return m * (m + uint64(1)) / uint64(2)
}

func SquareSequence(n int) uint64 {
	m := uint64(n)
	return m * m
}

func PentagonalSequence(n int) uint64 {
	m := uint64(n)
	return m * (uint64(3)*m - uint64(1)) / uint64(2)
}

func HexagonalSequence(n int) uint64 {
	m := uint64(n)
	return m * (uint64(2)*m - uint64(1))
}

func HeptagonalSequence(n int) uint64 {
	m := uint64(n)
	return m * (uint64(5)*m - uint64(3)) / uint64(2)
}

func OctagonalSequence(n int) uint64 {
	m := uint64(n)
	return m * (uint64(3)*m - uint64(2))
}
