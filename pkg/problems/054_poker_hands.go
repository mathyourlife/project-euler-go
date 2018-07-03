package problems

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PokerHands struct {
	pokerFile string
}

func (p *PokerHands) ID() int {
	return 54
}

func (p *PokerHands) Text() string {
	return `In the card game poker, a hand consists of five cards
and are ranked, from lowest to highest, in the following way:

    High Card:       Highest value card.
    One Pair:        Two cards of the same value.
    Two Pairs:       Two different pairs.
    Three of a Kind: Three cards of the same value.
    Straight:        All cards are consecutive values.
    Flush:           All cards of the same suit.
    Full House:      Three of a kind and a pair.
    Four of a Kind:  Four cards of the same value.
    Straight Flush:  All cards are consecutive values of same suit.
    Royal Flush:     Ten, Jack, Queen, King, Ace, in same suit.

The cards are valued in the order:
    2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.

If two players have the same ranked hands then the rank made up of the
highest value wins; for example, a pair of eights beats a pair of
fives (see example 1 below). But if two ranks tie, for example, both
players have a pair of queens, then highest cards in each hand are
compared (see example 4 below); if the highest cards tie then the next
highest cards are compared, and so on.

Consider the following five hands dealt to two players:
Hand    Player 1            Player 2           Winner
 1   5H 5C 6S 7S KD      2C 3S 8S 8D TD
      Pair of Fives       Pair of Eights       Player 2
 2   5D 8C 9S JS AC      2C 5C 7D 8S QH
      Highest card Ace    Highest card Queen   Player 1
 3   2D 9C AS AH AC      3D 6D 7D TD QD
      Three Aces          Flush with Diamonds  Player 2
 4   4D 6S 9H QH QC      3D 6D 7H QD QS
      Pair of Queens      Pair of Queens
      Highest card Nine   Highest card Seven   Player 1
 5   2H 2D 4C 4D 4S      3C 3D 3S 9S 9D
      Full House          Full House
      With Three Fours    with Three Threes    Player 1

The file, poker.txt, contains one-thousand random hands dealt to two
players. Each line of the file contains ten cards (separated by a
single space): the first five are Player 1's cards and the last five
are Player 2's cards. You can assume that all hands are valid (no
invalid characters or repeated cards), each player's hand is in no
specific order, and in each hand there is a clear winner.

How many hands does Player 1 win?

`
}

func (p *PokerHands) Solve() (string, error) {

	// var line string

	// line = "3S AD 9H JC 6D JD AS KH 6S JH"
	// line = "TS QS JD KS AS JD AS KH 6S JH" // royal
	// line = "TS QS JS KS AS JD AS KH 6S JH" // royal flush
	// line = "AS 2S 3S 4S 5S JD AS KH 6S JH" // straight flush
	// line = "3S 3D 3H 3C 6D JD AS KH 6S JH" // four of a kind
	// line = "3S 3S 3D JS JS JD AS KH 6S JH" // full house
	// line = "3S AS 9S JS 6S JD AS KH 6S JH" // flush
	// line = "AS 2S 3H 4S 5S JD AS KH 6S JH" // straight
	// line = "3S 3D 3H 2C 6D JD AS KH 6S JH" // three of a kind
	// line = "3S 3D 2H 2C 6D JD AS KH 6S JH" // two pairs
	// line = "3S 3D 2H AC 6D JD AS KH 6S JH" // pair
	// line = "7S 3D 2H AC 6D JD AS KH 6S JH" // pair

	// Pair of Fives       Pair of Eights       Player 2
	// line = "5H 5C 6S 7S KD 2C 3S 8S 8D TD"

	// Highest card Ace    Highest card Queen   Player 1
	// line = "5D 8C 9S JS AC 2C 5C 7D 8S QH"

	// Three Aces          Flush with Diamonds  Player 2
	// line = "2D 9C AS AH AC 3D 6D 7D TD QD"

	// Pair of Queens      Pair of Queens
	// Highest card Nine   Highest card Seven   Player 1
	// line = "4D 6S 9H QH QC 3D 6D 7H QD QS"

	// Full House          Full House
	// With Three Fours    with Three Threes    Player 1
	// line = "2H 2D 4C 4D 4S 3C 3D 3S 9S 9D"

	// match(line)
	// return "", nil

	file, err := os.Open(p.pokerFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	p1Wins := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if p.match(scanner.Text()) {
			p1Wins++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return fmt.Sprintf("%d", p1Wins), nil
}

func (p *PokerHands) match(line string) bool {
	cards := strings.Split(line, " ")
	l := cards[:5]
	r := cards[5:]
	p1Score := p.scoreHand(l)
	p2Score := p.scoreHand(r)
	for i := 0; i < len(p1Score); i++ {
		if p1Score[i] > p2Score[i] {
			// fmt.Println("player 1 wins")
			return true
		} else if p1Score[i] < p2Score[i] {
			// fmt.Println("player 2 wins")
			return false
		}
	}
	panic("tie")
	return false
}

func (p *PokerHands) cardValue(v byte) int {
	values := map[byte]int{
		[]byte("2")[0]: 2,
		[]byte("3")[0]: 3,
		[]byte("4")[0]: 4,
		[]byte("5")[0]: 5,
		[]byte("6")[0]: 6,
		[]byte("7")[0]: 7,
		[]byte("8")[0]: 8,
		[]byte("9")[0]: 9,
		[]byte("T")[0]: 10,
		[]byte("J")[0]: 11,
		[]byte("Q")[0]: 12,
		[]byte("K")[0]: 13,
		[]byte("A")[0]: 14,
	}
	return values[v]
}

func (p *PokerHands) scoreHand(hand []string) []int {
	if p.isRoyal(hand) && p.isFlush(hand) {
		// fmt.Println("royal flush")
		return []int{9}
	}
	if p.isStraight(hand) && p.isFlush(hand) {
		// fmt.Println("straight flush")
		return []int{8, p.highCard(hand)}
	}
	if ok, v, high := p.isOfAKind(hand, 4); ok {
		// fmt.Println("four of a kind")
		return []int{7, v, high}
	}
	if ok, high, low := p.isFullHouse(hand); ok {
		// fmt.Println("full house")
		return []int{6, high, low}
	}
	if p.isFlush(hand) {
		// fmt.Println("flush")
		return []int{5, p.highCard(hand)}
	}
	if p.isStraight(hand) {
		// fmt.Println("straight")
		return []int{4, p.highCard(hand)}
	}
	if ok, v, high := p.isOfAKind(hand, 3); ok {
		// fmt.Println("three of a kind")
		return []int{3, v, high}
	}
	if ok, high := p.isTwoPair(hand); ok {
		// fmt.Println("two pair")
		return []int{2, high}
	}
	if ok, v, high := p.isOfAKind(hand, 2); ok {
		// fmt.Println("pair")
		return []int{1, v, high}
	}
	// fmt.Println("high card")
	return []int{0, p.highCard(hand)}
}

func (p *PokerHands) highCard(hand []string) int {
	max := 0
	for _, card := range hand {
		v := p.cardValue(card[0])
		if v > max {
			max = v
		}
	}
	return max
}

func (p *PokerHands) isTwoPair(hand []string) (bool, int) {
	m := p.tallyValue(hand)
	if len(m) != 3 {
		return false, 0
	}
	countPairs := 0
	high := 0
	for v, count := range m {
		if count == 2 {
			countPairs++
		} else {
			high = v
		}
	}
	return countPairs == 2, high
}

func (p *PokerHands) isStraight(hand []string) bool {
	m := p.tallyValue(hand)
	if len(m) != 5 {
		return false
	}
	max := 0
	min := 0
	for _, card := range hand {
		v := p.cardValue(card[0])
		if v > max || max == 0 {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	if max-min == 4 {
		return true
	}
	if max != 14 || min != 2 {
		return false
	}
	// Check for ace low
	max = 0
	min = 0
	for _, card := range hand {
		v := p.cardValue(card[0])
		if v == 14 {
			v = 1
		}
		if v > max || max == 0 {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	return max-min == 4
}

func (p *PokerHands) isFlush(hand []string) bool {
	suit := byte(0)
	for _, card := range hand {
		if suit != byte(0) && suit != card[1] {
			return false
		}
		suit = card[1]
	}
	return true
}

func (p *PokerHands) isRoyal(hand []string) bool {
	r := map[int]bool{}
	for _, card := range hand {
		v := p.cardValue(card[0])
		if v >= 10 {
			r[v] = true
		}
	}
	return len(r) == 5
}

func (p *PokerHands) isFullHouse(hand []string) (bool, int, int) {
	m := p.tallyValue(hand)
	if len(m) != 2 {
		return false, 0, 0
	}
	low := 0
	high := 0
	for v, count := range m {
		if count == 2 {
			low = v
		} else if count == 3 {
			high = v
		}
	}
	return low > 0 && high > 0, high, low
}

func (p *PokerHands) isOfAKind(hand []string, n int) (bool, int, int) {
	m := p.tallyValue(hand)
	if len(m) != 6-n {
		return false, 0, 0
	}
	found := false
	high := 0
	kind := 0
	for v, count := range m {
		if count == n {
			found = true
			kind = v
		} else if v > high {
			high = v
		}
	}
	return found, kind, high
}

func (p *PokerHands) tallyValue(hand []string) map[int]int {
	m := map[int]int{}
	for _, card := range hand {
		m[p.cardValue(card[0])]++
	}
	return m
}
