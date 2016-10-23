/*
Poker hands

In the card game poker, a hand consists of five cards and are ranked,
from lowest to highest, in the following way:

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
*/
package main

import (
  "bufio"
	"log"
	"os"
)

const(
  HEART = 1 << iota
  DIAMOND
  CLUB
  SPADE
  CARD2
  CARD3
  CARD4
  CARD5
  CARD6
  CARD7
  CARD8
  CARD9
  CARDT
  CARDJ
  CARDQ
  CARDK
  CARDA
)

const (
  FACECARD = CARDT | CARDJ | CARDQ | CARDK | CARDA
  DIGIT = CARD2 | CARD3 | CARD4 | CARD5 | CARD6 | CARD7 | CARD8 | CARD9
  SUITS = HEART | DIAMOND | CLUB | SPADE
  VALUE = DIGIT | FACECARD
)

var ORDER = []int{CARDA, CARDK, CARDQ, CARDJ, CARDT, CARD9, CARD8, CARD7, CARD6, CARD5, CARD4, CARD3, CARD2}

var convert = map[byte]int{
  []byte("H")[0]: HEART,
  []byte("D")[0]: DIAMOND,
  []byte("C")[0]: CLUB,
  []byte("S")[0]: SPADE,
  []byte("2")[0]: CARD2,
  []byte("3")[0]: CARD3,
  []byte("4")[0]: CARD4,
  []byte("5")[0]: CARD5,
  []byte("6")[0]: CARD6,
  []byte("7")[0]: CARD7,
  []byte("8")[0]: CARD8,
  []byte("9")[0]: CARD9,
  []byte("T")[0]: CARDT,
  []byte("J")[0]: CARDJ,
  []byte("Q")[0]: CARDQ,
  []byte("K")[0]: CARDK,
  []byte("A")[0]: CARDA,
}

var convertBack = map[int]string{}

const (
  HighCard = 1 << iota
  OnePair
  TwoPairs
  ThreeOfAKind
  Straight
  Flush
  FullHouse
  FourOfAkind
  StraightFlush
  RoyalFlush
)

func parseHand(hand string) []int {
  parsedHand := []int{}
  for i := 0; i < 5; i++ {
    card := convert[hand[i*3+1]] | convert[hand[i*3]]
    parsedHand = append(parsedHand, card)
  }
  return parsedHand
}

func isRoyal(hand []int) bool {
  for _, c := range hand {
    if c & FACECARD == 0{
      return false
    }
  }
  return true
}

func isFlush(hand []int) bool {
  check := 0
  lastCheck := 0
  for _, c := range hand {
    check |= c & SUITS
    if lastCheck > 0 && lastCheck != check {
      return false
    }
    lastCheck = check
  }
  return true
}

func isStraight(hand []int) bool {
  cards := 0
  for _, c := range hand {
    cards |= c & VALUE
  }
  count := 0
  if cards & CARDA > 0 {
    count = 1
  }
  for {
    b := cards & 1
    cards = cards >> 1
    if b == 0 {
      count = 0
    } else {
      count++
      if count == 5 {
        return true
      }
    }
    if cards == 0 {
      break
    }
  }

  return false
}

func cardCount(hand []int) (int, int, map[int]int) {
  cc := map[int]int{}
  for _, c := range hand {
    cc[c & VALUE]++
  }
  maxCount := 0
  for _, v := range cc {
    if v > maxCount {
      maxCount = v
    }
  }
  maxCard := 0
  for k, v := range cc {
    if v == maxCount && k > maxCard {
      maxCard = k
    }
  }
  return maxCard, maxCount, cc
}

func isFullHouse(cc map[int]int) bool {
  if len(cc) != 2 {
    return false
  }
  for _, v := range cc {
    if v == 2 || v == 3 {
      return true
    } else {
      return false
    }
  }
  return false
}

func isTwoPairs(cc map[int]int) bool {
  if len(cc) != 3 {
    return false
  }
  count := 0
  for _, v := range cc {
    if v == 2 {
      count++
      if count == 2 {
        return true
      }
    }
  }
  return false
}

func scoreHand(hand []int) (int, int, int) {

  // HighCard
  // OnePair
  // TwoPairs
  // ThreeOfAKind
  // Straight
  // Flush
  // FullHouse
  // FourOfAkind
  // StraightFlush
  // RoyalFlush

  flush := isFlush(hand)
  royal := isRoyal(hand)
  if flush && royal {
    // log.Println("RoyalFlush")
    return RoyalFlush, 0, 0
  }

  straight := isStraight(hand)
  if flush && straight {
    // log.Println("StraightFlush")
    return StraightFlush, 0, 0
  }

  maxCard, maxCount, cc := cardCount(hand)
  if maxCount == 4 {
    // log.Println("FourOfAkind")
    return FourOfAkind, maxCard, 0
  }

  if isFullHouse(cc) {
    // log.Println("FullHouse")
    return FullHouse, 0, 0
  }

  if flush {
    // log.Println("Flush")
    return Flush, 0, 0
  }

  if straight {
    // log.Println("Straight")
    return Straight, 0, 0
  }

  if maxCount == 3 {
    // log.Println("ThreeOfAKind")
    return ThreeOfAKind, maxCard, 0
  }

  if isTwoPairs(cc) {
    // log.Println("TwoPairs")
    return TwoPairs, 0, 0
  }

  if maxCount == 2 {
    // log.Println("OnePair")
    v := 0
    for k, c := range hand {
      if k == maxCard {
        continue
      }
      v |= c & VALUE
    }
    return OnePair, maxCard, v
  }

  // log.Println("HighCard")
  v := 0
  for _, c := range hand {
    v |= c & VALUE
  }
  return HighCard, v, 0
}

func scoreGame(game string) int {

  // log.Println(game[:14])
  // log.Println(game[15:])
  hand1, h1, s1 := scoreHand(parseHand(game[:14]))
  hand2, h2, s2 := scoreHand(parseHand(game[15:]))
  // log.Println(hand1, convertBack[h1 & VALUE], hand2, convertBack[h2 & VALUE])

  if hand1 > hand2 {
    return 1
  } else if hand1 < hand2 {
    return 2
  }
  if h1 > h2 {
    return 1
  } else if h1 < h2 {
    return 2
  }
  if s1 > s2 {
    return 1
  } else if s1 < s2 {
    return 2
  }
  return 0
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Poker hands")

  for k, v := range convert {
    convertBack[v] = string(k)
  }

  file, err := os.Open("p054_poker.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  wins := map[int]int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    game := scanner.Text()
    winner := scoreGame(game)
    if winner == 0 {
      log.Println("UNKNOWN")
      log.Fatal("aoeu")
    }
    wins[winner]++
  }
  for player, count := range wins {
    log.Printf("Player %d wins: %d", player, count)
  }
  log.Printf("%#v\n", wins)
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  // handStr := "AD KD JD QD TD"  // RoyalFlush
  // handStr = "7S 6S 3S 4S 5S" // StraightFlush
  // handStr = "4C 4H 4D 5S 4S" // FourOfAkind
  // handStr = "5S 2H 5D 2D 2C" // FullHouse
  // handStr = "AS 3S 6S KS 4S" // Flush
  // handStr = "6D 3D 4D 5K 7S" // Straight
  // handStr = "9S KD QD 9H 9D" // ThreeOfAKind
  // handStr = "9S KD QD QH 9D" // ThreeOfAKind
  // handStr = "9S KD QD 4H 9D" // ThreeOfAKind
  // handStr = "9S KD QD 4H 2D" // ThreeOfAKind
  // log.Println(handStr)
  // hand := parseHand(handStr)
  // scoreHand(hand)

  // log.Println(isStraight(hand))
  // scoreHand(hand)


}