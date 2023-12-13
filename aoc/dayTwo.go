package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id       int
	games    []map[string]int
	minCubes map[string]int
}

func newGame() game {
	game := game{
		minCubes: map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		},
	}
	return game
}

var validSets = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGameData(gameData string) game {
	g := newGame()
	gameNumSplit := strings.Split(gameData, ": ")
	gameIDData := strings.Split(gameNumSplit[0], " ")
	gameID, _ := strconv.Atoi(gameIDData[1])
	g.id = gameID
	gameSets := strings.Split(gameNumSplit[1], "; ")
	for _, gameSet := range gameSets {
		g.games = append(g.games, createMapofSet(gameSet))
	}
	return g
}

func (g game) isGameValid() bool {
	for _, gameSet := range g.games {
		for key, val := range validSets {
			if gameSet[key] > val {
				return false
			}
		}
	}
	return true
}

func (g *game) setMinCubes() {
	for _, gameSet := range g.games {
		for key, val := range g.minCubes {
			if gameSet[key] > val {
				g.minCubes[key] = gameSet[key]
			}
		}
	}
}

func (g game) getPower() int {
	power := 1
	for _, val := range g.minCubes {
		power *= val
	}
	return power
}

func createMapofSet(puzzleInputSet string) map[string]int {
	cubeMap := make(map[string]int)
	cubeSet := strings.Split(puzzleInputSet, ", ")
	for _, c := range cubeSet {
		cubeData := strings.Split(c, " ")
		cubeType := cubeData[1]
		cubeVal, _ := strconv.Atoi(cubeData[0])
		cubeMap[cubeType] = cubeVal
	}
	return cubeMap
}

func DayTwoPartOne() {
	var sum int
	file, err := os.Open("./inputs/day_two_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		game := parseGameData(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if game.isGameValid() {
			sum += game.id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Puzzle Output: %d\n", sum)
}

func DayTwoPartTwo() {
	var sum int
	file, err := os.Open("./inputs/day_two_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		game := parseGameData(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		game.setMinCubes()
		sum += game.getPower()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Puzzle Output: %d\n", sum)
}
