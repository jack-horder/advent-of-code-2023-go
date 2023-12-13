package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getCalibrationValue(v string) (int, error) {
	chars := []rune(v)
	nums := make([]rune, 2)
	for _, r := range chars {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			nums[0] = r
			for i := len(chars) - 1; i >= 0; i-- {
				_, err := strconv.Atoi(string(chars[i]))
				if err == nil {
					nums[1] = chars[i]
					cb, err := strconv.Atoi(string(nums))
					if err != nil {
						return 0, err
					}
					return cb, nil
				}
			}
		}
	}
	return 0, nil
}

func replaceTextWords(v string) string {
	mutatedPuzzleInput := v
	textMap := map[string]string{
		"one":   "o1ne",
		"two":   "t2wo",
		"three": "t3hree",
		"four":  "f4our",
		"five":  "f5ive",
		"six":   "s6ix",
		"seven": "s7even",
		"eight": "e8ight",
		"nine":  "n9ine",
	}
	for key, val := range textMap {
		mutatedPuzzleInput = strings.ReplaceAll(mutatedPuzzleInput, key, val)
	}
	// log.Printf("Mutated Puzzle Input: %v", mutatedPuzzleInput)
	fmt.Println(mutatedPuzzleInput)
	return mutatedPuzzleInput
}

func DayOne() {
	file, err := os.Open("./inputs/day_one_input.txt")
	var puzzleOutput int
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		i, err := getCalibrationValue(replaceTextWords(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		puzzleOutput += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// i, _ := getCalibrationValue(replaceTextWords("two1nine"))
	// fmt.Printf("Puzzle Output: %d", i)
	fmt.Printf("Puzzle Output: %d\n", puzzleOutput)
}
