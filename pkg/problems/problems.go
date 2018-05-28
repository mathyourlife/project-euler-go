package problems

type Problem interface {
	ID() int
	Text() string
	Solve() (string, error)
}
