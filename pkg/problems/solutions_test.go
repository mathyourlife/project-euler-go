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
		{&PowerDigitSum{}, "1366"},
		{&NumberLetterCounts{}, "21124"},
		{&MaximumPathSumI{}, "1074"},
		{&CountingSundays{}, "171"},
		{&FactorialDigitSum{}, "648"},
		{&AmicableNumbers{}, "31626"},
		{&NameScores{namesFile: "testdata/p022_names.txt"}, "871198282"},
		{&NonAbundantSums{}, "4179871"},
		{&LexicographicPermutations{}, "2783915460"},
		{&ThousandDigitFibonacciNumber{}, "4782"},
		{&ReciprocalCycles{}, "983"},
		{&QuadraticPrimes{}, "-59231"},
		{&NumberSpiralDiagonals{}, "669171001"},
		{&DistinctPowers{}, "9183"},
		{&DigitFifthPowers{}, "443839"},
		{&CoinSums{}, "73682"},
		{&PandigitalProducts{}, "45228"},
		{&DigitCancellingFractions{}, "100"},
		{&DigitFactorials{}, "40730"},
		{&CircularPrimes{}, "55"},
		{&DoubleBasePalindromes{}, "872187"},
		{&TrunctablePrimes{}, "748317"},
		{&PandigitalMultiples{}, "932718654"},
		{&IntegerRightTriangles{}, "840"},
		{&ChampernowneConstant{}, "210"},
		{&PandigitalPrime{}, "7652413"},
		{&CodedTriangleNumbers{wordsFile: "testdata/p042_words.txt"}, "162"},
		{&SubStringDivisibility{}, "16695334890"},
		{&PentagonNumbers{}, "5482660"},
		{&TriangularPentagonalAndHexagonal{}, "1533776805"},
		{&GoldbachsOtherConjecture{}, "5777"},
		{&DistinctPrimesFactors{}, "134043"},
		{&SelfPowers{}, "9110846700"},
		// 296962999629
		// 997651
		// 121313
		// 142857
		// 4075
		// 376
	}

	found := false
	eulerProblem := os.Getenv("EULER_PROBLEM")

	for _, solution := range solutions {
		if eulerProblem != "" && fmt.Sprintf("%d", solution.Problem.ID()) != eulerProblem {
			continue
		}
		found = true

		t.Run(fmt.Sprintf("problem %d", solution.Problem.ID()), func(t *testing.T) {
			s, err := solution.Problem.Solve()
			if err != nil {
				t.Error(err)
			}
			if s != solution.Expected {
				t.Errorf("Problem: %d Expected solution: %s got: %s", solution.Problem.ID(), solution.Expected, s)
			}
		})
	}

	if !found && eulerProblem != "" {
		t.Errorf("unable to locate test case for problem %s", eulerProblem)
	}
}
