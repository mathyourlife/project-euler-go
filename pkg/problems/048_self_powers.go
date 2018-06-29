package problems

type SelfPowers struct{}

func (p *SelfPowers) ID() int {
	return 48
}

func (p *SelfPowers) Text() string {
	return `The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
`
}

func (p *SelfPowers) Solve() (string, error) {

	var m *BigInt
	s := NewBigInt(0)

	for n := 1; n <= 1000; n++ {
		// Calculate the self power
		m = NewBigInt(n)
		for i := 1; i < n; i++ {
			m.Mul(n)
			m.Regroup()
			// if the number is getting pretty big, we don't
			// need the most significant portion so drop down
			// to the 10 least significant digits
			if len(m.n) > 20 {
				m.n = m.n[:10]
			}
		}
		// Add it to the running sum
		s.AddBigInt(m)
		s.Regroup()
	}
	s.n = s.n[:10]

	return s.String(), nil
}
