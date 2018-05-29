package problems

import (
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
		// 70600674
	}

	for _, solution := range solutions {
		s, err := solution.Problem.Solve()
		if err != nil {
			t.Error(err)
		}
		if s != solution.Expected {
			t.Errorf("Problem: %d Expected solution: %s got: %s", solution.Problem.ID(), solution.Expected, s)
		}
	}
}
