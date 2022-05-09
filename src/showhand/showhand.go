package showhand

import (
	// "fmt"
	"sort"
	"constant"
)

type Card struct {
	rank int
	suit int
}

type Hand struct {
	card1 Card
	card2 Card
}

func compareCombination() bool {
	return false
}

func showHand(handCards []Card, boardCards []Card) (combination int) {
	
	return 0
}

func showCombination(cards [5]Card) (combination int, values [5]int) {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].rank < cards[j].rank
	})

	var count = [15]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		values[i] = cards[4-i].rank
		count[cards[i].rank]++
	}

	var three int = 0 
	var two1 int = 0
	var two2 int = 0
	for i := 2; i <= 14; i++ {
		if (count[i] == 4) {
			combination = constant.FOUROFAKIND
			return
		}

		if (count[i] == 3) {
			three = i
			continue
		}

		if (count[i] == 2) {
			if two1 == 0 {
				two1 = i
			} else {
				two2 = i
			}
		}
	}

	if three != 0 {
		if two1 != 0 {
			combination = constant.FULLHOUSE
		} else {
			combination = constant.THREEOFAKIND
		}

		return
	}

	if two1 != 0 {
		if two2 != 0 {
			combination = constant.TWOPAIRS
		} else {
			combination = constant.PAIR
		}

		return
	}
 
	var straight bool = checkStraight(cards)
	var flush bool = checkFlush(cards)

	if (straight && flush) {
		combination = constant.STRAIGHTFLUSH
		return
	}

	if (flush) {
		combination = constant.FLUSH
		return
	}

	if (straight) {
		combination = constant.STRAIGHT
		return
	}

	combination = constant.HIGHCARD
	return
}

func checkStraight(cards [5]Card) bool {
	for i := 1; i < 5; i++ {
		if (cards[i].rank != cards[i-1].rank + 1) {
			return false
		}
	}
	return true
}

func checkFlush(cards [5]Card) bool {
	if (cards[0].suit == cards[1].suit && cards[1].suit == cards[2].suit && cards[2].suit == cards[3].suit && cards[3].suit == cards[4].suit) {
		return true
	}
	return false
}

func printResult() int {

	return 0
}