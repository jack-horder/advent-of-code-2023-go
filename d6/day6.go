package d6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time      int
	dist      int
	waysToWin int
}

func readRaces(filename string) ([]string, error) {
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

func (r *race) setWaysToWin() {
	for i := 0; i < r.time; i++ {
		remaining := r.time - i
		remaining *= i
		if remaining > r.dist {
			r.waysToWin++
		}
	}
}

func parseRaces(lines []string) ([]*race, error) {
	var races []*race
	times := strings.Fields(strings.Split(lines[0], ": ")[1])
	distances := strings.Fields(strings.Split(lines[1], ": ")[1])
	for i := 0; i < len(times); i++ {
		timeVal, err := strconv.Atoi(times[i])
		if err != nil {
			return nil, err
		}
		distVal, err := strconv.Atoi(distances[i])
		if err != nil {
			return nil, err
		}
		race := &race{
			time: timeVal,
			dist: distVal,
		}
		races = append(races, race)
	}
	return races, nil
}

func parseRacesPartTwo(lines []string) ([]*race, error) {
	var races []*race
	time := strings.Join(strings.Fields(strings.Split(lines[0], ": ")[1]), "")
	distance := strings.Join(strings.Fields(strings.Split(lines[1], ": ")[1]), "")
	timeVal, err := strconv.Atoi(time)
	if err != nil {
		return nil, err
	}
	distVal, err := strconv.Atoi(distance)
	if err != nil {
		return nil, err
	}
	race := &race{
		time: timeVal,
		dist: distVal,
	}
	races = append(races, race)
	return races, nil
}

func DaySixPartOne() {
	lines, err := readRaces("./inputs/6_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	races, err := parseRaces(lines)
	if err != nil {
		log.Fatalln(err)
	}
	answer := 1
	for _, race := range races {
		race.setWaysToWin()
		answer *= race.waysToWin
	}
	fmt.Printf("Puzzle Output: %d\n", answer)
}

func DaySixPartTwo() {
	lines, err := readRaces("./inputs/6_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	races, err := parseRacesPartTwo(lines)
	if err != nil {
		log.Fatalln(err)
	}
	answer := 1
	for _, race := range races {
		race.setWaysToWin()
		answer *= race.waysToWin
	}
	fmt.Printf("Puzzle Output: %d\n", answer)
}
