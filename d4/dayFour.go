package d4

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type card struct {
	cardNum          int
	winningNumbers   []int
	chosenNumbers    []int
	winningNumsCount int
	score            int
}

func (c *card) setScore() {
	winningNumsCount := 0
	for _, winNums := range c.winningNumbers {
		for _, choseNums := range c.chosenNumbers {
			if winNums == choseNums {
				winningNumsCount += 1
			}
		}
	}
	c.winningNumsCount = winningNumsCount
	c.score = int(math.Pow(2, float64(winningNumsCount)-1))
}

func getIntSlice(strNums []string) []int {
	var intSlice []int
	for _, i := range strNums {
		num, _ := strconv.Atoi(i)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func createCard(cardData string) *card {
	c := &card{}
	removedExtraSpaces := strings.ReplaceAll(cardData, "   ", " ")
	removedExtraSpaces = strings.ReplaceAll(removedExtraSpaces, "  ", " ")
	dataSplitCard := strings.Split(removedExtraSpaces, ": ")
	cardMetaData := dataSplitCard[0]
	cardNum, _ := strconv.Atoi(strings.Split(cardMetaData, " ")[1])
	c.cardNum = cardNum
	numbersSplit := strings.Split(dataSplitCard[1], " | ")
	winningNumbers := strings.Split(numbersSplit[0], " ")
	chosenNumbers := strings.Split(numbersSplit[1], " ")
	c.winningNumbers = getIntSlice(winningNumbers)
	c.chosenNumbers = getIntSlice(chosenNumbers)
	c.setScore()
	return c
}

func parsePuzzleInput(filename string) []*card {
	var cards []*card
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, createCard(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cards
}

func checkWins(origCards []*card, newCards []*card) []*card {
	var cards []*card
	for _, card := range newCards {
		if card.winningNumsCount > 0 {
			if card.cardNum != len(origCards) {
				cardIdx := card.cardNum
				toCardIdx := cardIdx + card.winningNumsCount
				if toCardIdx == len(origCards)+1 {
					toCardIdx = len(origCards)
				}
				copyCards := origCards[cardIdx:toCardIdx]
				cards = append(cards, copyCards...)
			}
		}
	}
	return cards
}

func DayFourPartOne() {
	var score int
	cards := parsePuzzleInput("./inputs/4_input.txt")
	for _, c := range cards {
		score += c.score
	}
	fmt.Printf("Puzzle Output: %d\n", score)
	// fmt.Printf("Card Number: %d\nWinning Numbers: %v\nChosen Numbers: %v\nScore: %d\n", card.cardNum, card.winningNumbers, card.chosenNumbers, card.score)
}

func DayFourPartTwo() {
	var totalCards int
	cards := parsePuzzleInput("./inputs/4_input.txt")
	newCards := checkWins(cards, cards)
	totalCards += len(cards)
	for len(newCards) != 0 {
		totalCards += len(newCards)
		newCards = checkWins(cards, newCards)
	}
	fmt.Printf("Puzzle Output: %d\n", totalCards)
	// fmt.Printf("Card Number: %d\nWinning Numbers: %v\nChosen Numbers: %v\nScore: %d\n", cards[0].cardNum, cards[0].winningNumbers, cards[0].chosenNumbers, cards[0].score)

}
