package d9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reportLine struct {
	sequence       []int
	nestedSeq      []*reportLine
	prediction     int
	backPrediction int
}

func (r reportLine) String() string {
	var allLines [][]int
	allLines = append(allLines, r.sequence)
	for _, seq := range r.nestedSeq {
		allLines = append(allLines, seq.sequence)
	}
	return fmt.Sprint(allLines)
}

func createSequence(r *reportLine) *reportLine {
	rl := &reportLine{}
	for i := 0; i < len(r.sequence); i++ {
		if i != (len(r.sequence) - 1) {
			diff := r.sequence[i+1] - r.sequence[i]
			rl.sequence = append(rl.sequence, diff)
		}
	}
	return rl
}

func (r *reportLine) setNestedSequences() {
	var reportLines []*reportLine
	rl := r
	for !checkRepSeq(rl) {
		rl = createSequence(rl)
		reportLines = append(reportLines, rl)
	}
	r.nestedSeq = reportLines
}

func (r *reportLine) calculatePredictions() {
	sum := 0
	for i := 0; i < len(r.nestedSeq); i++ {
		cur := r.nestedSeq[i]
		sum += cur.sequence[len(cur.sequence)-1]
	}
	r.prediction = r.sequence[len(r.sequence)-1] + sum
}

func (r *reportLine) calculateBackPredictions() {
	sum := 0
	for i := len(r.nestedSeq) - 1; i > 0; i-- {
		if i == len(r.nestedSeq)-1 {
			cur := r.nestedSeq[i]
			prev := r.nestedSeq[i-1]
			sum = prev.sequence[0] - cur.sequence[0]
			continue
		}
		cur := r.nestedSeq[i-1]
		sum = cur.sequence[0] - sum
	}
	r.backPrediction = r.sequence[0] - sum
}

func readLines(filename string) ([]string, error) {
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

func parseReportLines(lines []string) []*reportLine {
	repLines := []*reportLine{}
	for _, line := range lines {
		rl := new(reportLine)
		for _, n := range strings.Fields(line) {
			val, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalln(err)
			}
			rl.sequence = append(rl.sequence, val)
		}
		repLines = append(repLines, rl)
	}
	return repLines
}

func checkRepSeq(rl *reportLine) bool {
	allZero := true
	for _, i := range rl.sequence {
		allZero = allZero && (i == 0)
	}
	return allZero
}

func DayNinePartOne() {
	lines, err := readLines("./inputs/9_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	repLines := parseReportLines(lines)
	for _, r := range repLines {
		r.setNestedSequences()
	}
	sum := 0
	for _, rl := range repLines {
		rl.calculatePredictions()
		fmt.Printf("report line: %v, prediction: %d\n", rl.sequence, rl.prediction)
		sum += rl.prediction
	}
	fmt.Printf("Puzzle Output: %d\n", sum)
}

func DayNinePartTwo() {
	lines, err := readLines("./inputs/9_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	repLines := parseReportLines(lines)
	for _, r := range repLines {
		r.setNestedSequences()
	}
	sum := 0
	for _, rl := range repLines {
		rl.calculateBackPredictions()
		fmt.Printf("report line: %v, prediction: %d\n", rl.sequence, rl.backPrediction)
		sum += rl.backPrediction
	}
	fmt.Printf("Puzzle Output: %d\n", sum)
}
