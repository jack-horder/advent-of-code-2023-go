package d3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type partPos struct {
	row int
	col int
}

type part struct {
	partNum          int
	startPos         partPos
	endPos           partPos
	isSymbolAdjacent bool
	symbolRune       rune
	symbolPos        partPos
	checked          bool
}

func getStarParts(parts []*part) []*part {
	var gearParts []*part
	for _, part := range parts {
		if string(part.symbolRune) == "*" {
			gearParts = append(gearParts, part)
		}
	}
	return gearParts
}

func (p *part) setIsSymbolAdjacent(puzzleInput [][]rune) {
	// check top
	if p.startPos.row != 0 {
		for i := p.startPos.col; i <= p.endPos.col; i++ {
			r := puzzleInput[p.startPos.row-1][i]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.startPos.row - 1, col: i}
					return
				}
			}
		}
	}
	// check bottom
	if p.startPos.row != len(puzzleInput)-1 {
		for i := p.startPos.col; i <= p.endPos.col; i++ {
			r := puzzleInput[p.startPos.row+1][i]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.startPos.row + 1, col: i}
					return
				}
			}
		}
	}
	// check left
	if p.startPos.col != 0 {
		r := puzzleInput[p.startPos.row][p.startPos.col-1]
		if string(r) != "." {
			if !unicode.IsNumber(r) {
				p.isSymbolAdjacent = true
				p.symbolRune = r
				p.symbolPos = partPos{row: p.startPos.row, col: p.startPos.col - 1}
				return
			}
		}
	}
	// check right
	if p.endPos.col != len(puzzleInput[p.endPos.row])-1 {
		r := puzzleInput[p.endPos.row][p.endPos.col+1]
		if string(r) != "." {
			if !unicode.IsNumber(r) {
				p.isSymbolAdjacent = true
				p.symbolRune = r
				p.symbolPos = partPos{row: p.endPos.row, col: p.endPos.col + 1}
				return
			}
		}
	}
	// check top left diagonal
	if p.startPos.col != 0 {
		if p.startPos.row != 0 {
			r := puzzleInput[p.startPos.row-1][p.startPos.col-1]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.startPos.row - 1, col: p.startPos.col - 1}
					return
				}
			}
		}
	}
	// check bottom left diagonal
	if p.startPos.col != 0 {
		if p.startPos.row != len(puzzleInput)-1 {
			r := puzzleInput[p.startPos.row+1][p.startPos.col-1]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.startPos.row + 1, col: p.startPos.col - 1}
					return
				}
			}
		}
	}
	// check top right diagonal
	if p.endPos.col != len(puzzleInput[p.endPos.row])-1 {
		if p.endPos.row != 0 {
			r := puzzleInput[p.endPos.row-1][p.endPos.col+1]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.endPos.row - 1, col: p.endPos.col + 1}
					return
				}
			}
		}
	}
	// check bottom right diagonal
	if p.endPos.col != len(puzzleInput[p.endPos.row])-1 {
		if p.endPos.row != len(puzzleInput)-1 {
			r := puzzleInput[p.endPos.row+1][p.endPos.col+1]
			if string(r) != "." {
				if !unicode.IsNumber(r) {
					p.isSymbolAdjacent = true
					p.symbolRune = r
					p.symbolPos = partPos{row: p.endPos.row + 1, col: p.endPos.col + 1}
					return
				}
			}
		}
	}
}

func loadInput(filename string) [][]rune {
	var puzzleInput [][]rune
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
		row := []rune(scanner.Text())
		puzzleInput = append(puzzleInput, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return puzzleInput
}

func getParts(puzzleInput [][]rune) []*part {
	var parts []*part
	for rowIdx, row := range puzzleInput {
		p := &part{}
		partNum := []rune{}
		foundPartNum := false
		for colIdx, col := range row {
			if unicode.IsNumber(col) {
				partNum = append(partNum, col)
				if !foundPartNum {
					foundPartNum = true
					p.startPos = partPos{row: rowIdx, col: colIdx}
				}
				continue
			}
			if foundPartNum {
				p.endPos = partPos{row: rowIdx, col: colIdx - 1}
				partNumInt, _ := strconv.Atoi(string(partNum))
				p.partNum = partNumInt
				parts = append(parts, p)
				partNum = []rune{}
				p = &part{}
				foundPartNum = false
			}
		}
		if foundPartNum {
			p.endPos = partPos{row: rowIdx, col: len(row) - 1}
			partNumInt, _ := strconv.Atoi(string(partNum))
			p.partNum = partNumInt
			parts = append(parts, p)
			p = &part{}
			foundPartNum = false
		}
	}
	return parts
}

func DayThreePartOne() {
	data := loadInput("./inputs/3_input.txt")
	parts := getParts(data)
	for _, part := range parts {
		part.setIsSymbolAdjacent(data)
		fmt.Printf("Part Number: %d, Has Adjacent Symbol: %v\n", part.partNum, part.isSymbolAdjacent)
	}
	var sum int
	for _, part := range parts {
		if part.isSymbolAdjacent {
			sum += part.partNum
		}
	}
	fmt.Println("Total Rows: " + fmt.Sprint(len(data)))
	fmt.Println("Total Part Numbers: " + fmt.Sprint(len(parts)))
	fmt.Println("Sum of Part Numbers: " + fmt.Sprint(sum))
}

func DayThreePartTwo() {
	data := loadInput("./inputs/3_input_harry.txt")
	parts := getParts(data)
	for _, part := range parts {
		part.setIsSymbolAdjacent(data)
		fmt.Printf("Part Number: %d, Has Adjacent Symbol: %v\n", part.partNum, part.isSymbolAdjacent)
	}
	var sum int

	starParts := getStarParts(parts)
	for _, part := range starParts {
		for _, compPart := range starParts {
			if part != compPart {
				if part.symbolPos == compPart.symbolPos {
					if !compPart.checked {
						sum = sum + (part.partNum * compPart.partNum)
						part.checked = true
						compPart.checked = true
						break
					}
				}
			}
		}
	}
	fmt.Println("Sum of Gears: " + fmt.Sprint(sum))
}
