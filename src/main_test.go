package main

import (
	"testing"
)

func TestPokerHandRank(t *testing.T) {
	tests := []struct {
		name     string
		cards    []string
		expected int
	}{
		{
			name:     "HighCard",
			cards:    []string{"2D", "4H", "6C", "8S", "QH"},
			expected: HighCard,
		},
		{
			name:     "OnePair",
			cards:    []string{"3D", "3H", "5C", "8S", "QH"},
			expected: OnePair,
		},
		{
			name:     "TwoPairs",
			cards:    []string{"3D", "3H", "5C", "5S", "QH"},
			expected: TwoPairs,
		},
		{
			name:     "ThreeOfAKind",
			cards:    []string{"3D", "3H", "3C", "5S", "QH"},
			expected: ThreeOfAKind,
		},
		{
			name:     "Straight",
			cards:    []string{"8D", "9H", "10C", "JS", "QH"},
			expected: Straight,
		},
		// Add more test cases here...
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PokerHandRank(test.cards)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", handRanks[test.expected], handRanks[result])
			}
		})
	}
}
