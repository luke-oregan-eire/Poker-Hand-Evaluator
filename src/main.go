package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FourOfAKind
	Straight
	Flush
	FullHouse
)

var handRanks = map[int]string{
	HighCard:     "High Card",
	OnePair:      "One Pair",
	TwoPairs:     "Two Pairs",
	ThreeOfAKind: "Three of a Kind",
	FourOfAKind:  "Four of a Kind",
	Straight:     "Straight",
	Flush:        "Flush",
	FullHouse:    "Full House",
}

func PokerHandRank(cards []string) int {
	ranks := make(map[string]int)
	suits := make(map[string]rune)

	for _, card := range cards {
		rank := card[:len(card)-1]
		suit := rune(card[len(card)-1])
		ranks[rank]++
		suits[rank] = suit
	}

	pairCount := 0
	threeCount := 0

	for _, count := range ranks {
		switch count {
		case 2:
			pairCount++
		case 3:
			threeCount++
		}
	}

	if len(ranks) == 2 {
		if pairCount == 1 && threeCount == 1 {
			return FullHouse
		} else if pairCount == 2 {
			return TwoPairs
		}
	} else if len(ranks) == 3 {
		if pairCount == 1 && threeCount == 1 {
			return FullHouse
		}
	} else if len(ranks) == 4 {
		return OnePair
	} else if len(ranks) == 5 {
		if isStraight(cards) {
			return Straight
		} else if len(suits) == 1 {
			return Flush
		}
	}

	return HighCard
}

func isStraight(cards []string) bool {
	rankOrder := "23456789TJQKA"
	rankSet := make(map[byte]bool)
	for _, card := range cards {
		rankSet[card[0]] = true
	}

	// Check for Ace-low straight (A-5-4-3-2)
	if len(rankSet) == 5 && rankSet['A'] && rankSet['2'] && rankSet['3'] && rankSet['4'] && rankSet['5'] {
		return true
	}

	// Check for other straights
	for i := 0; i <= len(rankOrder)-5; i++ {
		if rankSet[rankOrder[i]] && rankSet[rankOrder[i+1]] && rankSet[rankOrder[i+2]] &&
			rankSet[rankOrder[i+3]] && rankSet[rankOrder[i+4]] {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter five cards (e.g., 10D QD 5H KS 3C): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	cardStrings := strings.Split(input, " ")

	if len(cardStrings) != 5 {
		fmt.Println("Please provide exactly five cards.")
		return
	}

	cards := make([]string, 0, 5)
	for _, cardStr := range cardStrings {
		cards = append(cards, cardStr)
	}

	rank := PokerHandRank(cards)
	fmt.Printf("Best Poker Hand: %s\n", handRanks[rank])
}
