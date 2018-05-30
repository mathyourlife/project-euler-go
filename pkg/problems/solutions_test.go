package problems

import (
	"fmt"
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	solutions := []struct {
		Problem  Problem
		Expected string
	}{
		{&MultiplesOf3Or5{}, "233168"},
		{&EvenFibonacciNumbers{}, "4613732"},
		{&LargestPrimeFactor{}, "6857"},
		{&LargestPalindromeProduct{}, "906609"},
		{&SmallestMultiple{}, "232792560"},
		{&SumSquareDifference{}, "25164150"},
		{&TenThousandFirstPrime{}, "104743"},
		{&LargestProductInASeries{}, "23514624000"},
		{&SpecialPythagoreanTriplet{}, "31875000"},
		{&SummationOfPrimes{}, "142913828922"},
		{&LargestProductInAGrid{}, "70600674"},
		{&HighlyDivisibleTriangularNumber{}, "76576500"},
		{&LargeSum{}, "5537376230"},
		{&LongestCollatzSequence{}, "837799"},
		{&LatticePaths{}, "137846528820"},
		// 1366
		// 21124
		// 1074
		// 171
		// 648
		// 31626
		// 871198282
		// 4179871
		// 2783915460
	}

	eulerProblem := os.Getenv("EULER_PROBLEM")

	for _, solution := range solutions {
		if eulerProblem != "" && fmt.Sprintf("%d", solution.Problem.ID()) != eulerProblem {
			continue
		}
		s, err := solution.Problem.Solve()
		if err != nil {
			t.Error(err)
		}
		if s != solution.Expected {
			t.Errorf("Problem: %d Expected solution: %s got: %s", solution.Problem.ID(), solution.Expected, s)
		}
	}
}
