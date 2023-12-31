package d7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	val             string
	numVal          int
	numOfOccurences int
}

type hand struct {
	cards        []card
	handMap      map[string]int
	handVals     []int
	handValsStr  []string
	bid          int
	highCardVal  int
	handStrength int
}

func (h hand) cardsString() string {
	return strings.Join(h.handValsStr, ", ")
}

func (h *hand) setHandStrength() {
	// give score of 1-7 based on type of hand
	// check for 5 of a kind
	cards := sortCards(h.cards)
	if cards[0].numOfOccurences == 5 {
		h.handStrength = 7
		return
	}
	// check for 4 of a kind
	if cards[0].numOfOccurences == 4 {
		h.handStrength = 6
		return
	}
	// full house
	if cards[0].numOfOccurences == 3 {
		if cards[1].numOfOccurences == 2 {
			h.handStrength = 5
			return
		}
	}
	// 3 of a kind
	if cards[0].numOfOccurences == 3 {
		h.handStrength = 4
		return
	}
	// two pair
	if cards[0].numOfOccurences == 2 {
		if cards[1].numOfOccurences == 2 {
			h.handStrength = 3
			return
		}
	}
	// one pair
	if cards[0].numOfOccurences == 2 {
		h.handStrength = 2
		return
	}
	// high card
	h.handStrength = 1
}

func createCardMap() map[string]int {
	cards := []string{
		"A",
		"K",
		"Q",
		"J",
		"T",
		"9",
		"8",
		"7",
		"6",
		"5",
		"4",
		"3",
		"2",
	}
	cardMap := make(map[string]int)
	for i, k := 0, len(cards)+1; i < len(cards); i, k = i+1, k-1 {
		cardMap[cards[i]] = k
	}
	return cardMap
}

func createCardMapPartTwo() map[string]int {
	cards := []string{
		"A",
		"K",
		"Q",
		"T",
		"9",
		"8",
		"7",
		"6",
		"5",
		"4",
		"3",
		"2",
		"J",
	}
	cardMap := make(map[string]int)
	for i, k := 0, len(cards)+1; i < len(cards); i, k = i+1, k-1 {
		cardMap[cards[i]] = k
	}
	return cardMap
}

func sortCards(cards []card) []card {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].numOfOccurences > cards[j].numOfOccurences
	})
	return cards
}

func repurposeJokers(cards []card) []card {
	var redistCards []card
	var jokerCard card
	for _, c := range cards {
		if c.val == "J" {
			jokerCard = c
			continue
		}
		redistCards = append(redistCards, c)
	}
	if (card{}) == jokerCard {
		return cards
	}
	if len(redistCards) == 0 {
		return cards
	}
	redistCards[0].numOfOccurences += jokerCard.numOfOccurences
	return redistCards
}

func parseHand(handStr string, cardMap map[string]int) *hand {
	h := new(hand)
	cu := make(map[string]int)
	cards := []rune(strings.Fields(handStr)[0])
	bid, _ := strconv.Atoi(strings.Fields(handStr)[1])
	for _, card := range cards {
		val, ok := cu[string(card)]
		if ok {
			cu[string(card)] = val + 1
			continue
		}
		cu[string(card)] = 1
	}
	h.bid = bid
	for key, val := range cu {
		numVal := cardMap[key]
		c := card{
			val:             key,
			numVal:          numVal,
			numOfOccurences: val,
		}
		h.cards = append(h.cards, c)
	}
	h.cards = sortCards(h.cards)
	handVals := []string{}
	for _, r := range cards {
		val := cardMap[string(r)]
		handVals = append(handVals, fmt.Sprintf("%v", val))
	}
	handValInts := []int{}
	for _, card := range handVals {
		val, _ := strconv.Atoi(card)
		handValInts = append(handValInts, val)
	}
	h.handMap = cu
	h.handVals = handValInts
	h.handValsStr = handVals
	h.highCardVal = slices.Max(handValInts)
	return h
}

func readHands(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func DaySevenPartOne() {
	handsStr, err := readHands("./inputs/7_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	m := createCardMap()
	var hands []*hand
	for _, hand := range handsStr {
		h := parseHand(hand, m)
		h.setHandStrength()
		hands = append(hands, h)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handStrength == hands[j].handStrength {
			for x := range hands[i].handVals {
				if hands[i].handVals[x] == hands[j].handVals[x] {
					continue
				}
				return hands[i].handVals[x] < hands[j].handVals[x]
			}
		}
		return hands[i].handStrength < hands[j].handStrength
	})
	score := 0
	for i, hand := range hands {
		fmt.Printf("Cards: %v, Hand Rank: %d, Bid: %d, Total Winnings: %d\n", hand.cardsString(), hand.handStrength, hand.bid, (hand.bid * (i + 1)))
		score += (hand.bid * (i + 1))
	}
	fmt.Printf("Puzzle Output: %d\n", score)
}

func DaySevenPartTwo() {
	handsStr, err := readHands("./inputs/7_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	m := createCardMapPartTwo()
	var hands []*hand
	for _, hand := range handsStr {
		h := parseHand(hand, m)
		h.cards = repurposeJokers(h.cards)
		h.setHandStrength()
		hands = append(hands, h)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handStrength == hands[j].handStrength {
			for x := range hands[i].handVals {
				if hands[i].handVals[x] == hands[j].handVals[x] {
					continue
				}
				return hands[i].handVals[x] < hands[j].handVals[x]
			}
		}
		return hands[i].handStrength < hands[j].handStrength
	})
	score := 0
	for i, hand := range hands {
		fmt.Printf("Cards: %v, Hand Rank: %d, Bid: %d, Total Winnings: %d\n", hand.cardsString(), hand.handStrength, hand.bid, (hand.bid * (i + 1)))
		score += (hand.bid * (i + 1))
	}
	fmt.Printf("Puzzle Output: %d\n", score)
}
