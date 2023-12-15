package d5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type seed struct {
	seedNum       int
	soilNum       int
	fertilizerNum int
	waterNum      int
	lightNum      int
	tempNum       int
	humidityNum   int
	locationNum   int
}

type seedRange struct {
	seedStart     int
	seedRange     int
	soilNum       int
	fertilizerNum int
	waterNum      int
	lightNum      int
	tempNum       int
	humidityNum   int
	locationNum   int
}

type mapRanges []*mapRange

type mapRange struct {
	srcRangeStart  int
	srcRangeEnd    int
	destRangeStart int
}

func newMapRange(srcRangeStart int, rangeLen int, destRangeStart int) *mapRange {
	m := &mapRange{
		srcRangeStart:  srcRangeStart,
		srcRangeEnd:    srcRangeStart + rangeLen,
		destRangeStart: destRangeStart,
	}
	return m
}

func (mr mapRanges) getMapValue(srcVal int) int {
	for _, m := range mr {
		mapVal, exists := m.getDestVal(srcVal)
		if exists {
			return mapVal
		}
	}
	return srcVal
}

func (m mapRange) getDestVal(srcVal int) (int, bool) {
	if m.srcRangeStart <= srcVal && srcVal < m.srcRangeEnd {
		return (m.destRangeStart + (srcVal - m.srcRangeStart)), true
	}
	return 0, false
}

func (s *seed) setMapNums(
	soilMap mapRanges,
	fertilizerMap mapRanges,
	waterMap mapRanges,
	lightMap mapRanges,
	tempMap mapRanges,
	humdidityMap mapRanges,
	locationMap mapRanges,
) {
	// soil map
	s.soilNum = soilMap.getMapValue(s.seedNum)
	// fertilizer map
	s.fertilizerNum = fertilizerMap.getMapValue(s.soilNum)
	// water map
	s.waterNum = waterMap.getMapValue(s.fertilizerNum)
	// light map
	s.lightNum = lightMap.getMapValue(s.waterNum)
	// temp map
	s.tempNum = tempMap.getMapValue(s.lightNum)
	// humidity map
	s.humidityNum = humdidityMap.getMapValue(s.tempNum)
	// location map
	s.locationNum = locationMap.getMapValue(s.humidityNum)
}

func (s *seedRange) setMapNums(
	soilMap mapRanges,
	fertilizerMap mapRanges,
	waterMap mapRanges,
	lightMap mapRanges,
	tempMap mapRanges,
	humdidityMap mapRanges,
	locationMap mapRanges,
) {
	var soilNum, fertilizerNum, waterNum, lightNum, tempNum, humidityNum, locationNum int
	for i := s.seedStart; i < s.seedStart+s.seedRange; i++ {
		// soil map
		soilNum = soilMap.getMapValue(i)
		// fertilizer map
		fertilizerNum = fertilizerMap.getMapValue(soilNum)
		// water map
		waterNum = waterMap.getMapValue(fertilizerNum)
		// light map
		lightNum = lightMap.getMapValue(waterNum)
		// temp map
		tempNum = tempMap.getMapValue(lightNum)
		// humidity map
		humidityNum = humdidityMap.getMapValue(tempNum)
		// location map
		locationNum = locationMap.getMapValue(humidityNum)
		if s.locationNum == 0 || locationNum < s.locationNum {
			s.soilNum = soilNum
			s.fertilizerNum = fertilizerNum
			s.waterNum = waterNum
			s.lightNum = lightNum
			s.tempNum = tempNum
			s.humidityNum = humidityNum
			s.locationNum = locationNum
			continue
		}
	}

}

func parseSeeds(seedInput string) []*seed {
	var seeds []*seed
	seedsStr := strings.Split(strings.Split(seedInput, ": ")[1], " ")
	for _, seedNum := range seedsStr {
		seedInt, _ := strconv.Atoi(seedNum)
		s := &seed{seedNum: seedInt}
		seeds = append(seeds, s)
	}
	return seeds
}

func parseSeedsPart2(seedInput string) []*seedRange {
	var seeds []*seedRange
	seedsStr := strings.Split(strings.Split(seedInput, ": ")[1], " ")
	for i := 0; i < len(seedsStr); i += 2 {
		ss, _ := strconv.Atoi(seedsStr[i])
		sr, _ := strconv.Atoi(seedsStr[i+1])
		s := &seedRange{seedStart: ss, seedRange: sr}
		seeds = append(seeds, s)
	}
	return seeds
}

func createMap(mapRows [][]int) mapRanges {
	mr := mapRanges{}
	for _, row := range mapRows {
		m := newMapRange(
			row[1],
			row[2],
			row[0],
		)
		mr = append(mr, m)
	}
	return mr
}

func getMapSlice(s *bufio.Scanner) [][]int {
	var strNums [][]string
	var intNums [][]int
	for s.Scan() {
		if len(s.Text()) == 0 {
			for _, numArr := range strNums {
				var ints []int
				for _, num := range numArr {
					intNum, _ := strconv.Atoi(num)
					ints = append(ints, intNum)
				}
				intNums = append(intNums, ints)
			}
			return intNums
		}
		if len(s.Text()) != 0 {
			if !strings.Contains(s.Text(), ":") {
				strNums = append(strNums, strings.Split(s.Text(), " "))
			}
		}
	}
	for _, numArr := range strNums {
		var ints []int
		for _, num := range numArr {
			intNum, _ := strconv.Atoi(num)
			ints = append(ints, intNum)
		}
		intNums = append(intNums, ints)
	}
	return intNums
}

func parseInput(filename string) (
	seeds []*seed,
	soilMap mapRanges,
	fertilizerMap mapRanges,
	waterMap mapRanges,
	lightMap mapRanges,
	tempMap mapRanges,
	humdidityMap mapRanges,
	locationMap mapRanges,
) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	scanner.Scan()
	seeds = parseSeeds(scanner.Text())
	scanner.Scan()
	soilMap = createMap(getMapSlice(scanner))
	fertilizerMap = createMap(getMapSlice(scanner))
	waterMap = createMap(getMapSlice(scanner))
	lightMap = createMap(getMapSlice(scanner))
	tempMap = createMap(getMapSlice(scanner))
	humdidityMap = createMap(getMapSlice(scanner))
	locationMap = createMap(getMapSlice(scanner))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func parseInputPart2(filename string) (
	seeds []*seedRange,
	soilMap mapRanges,
	fertilizerMap mapRanges,
	waterMap mapRanges,
	lightMap mapRanges,
	tempMap mapRanges,
	humdidityMap mapRanges,
	locationMap mapRanges,
) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	scanner.Scan()
	seeds = parseSeedsPart2(scanner.Text())
	scanner.Scan()
	soilMap = createMap(getMapSlice(scanner))
	fertilizerMap = createMap(getMapSlice(scanner))
	waterMap = createMap(getMapSlice(scanner))
	lightMap = createMap(getMapSlice(scanner))
	tempMap = createMap(getMapSlice(scanner))
	humdidityMap = createMap(getMapSlice(scanner))
	locationMap = createMap(getMapSlice(scanner))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func calculateMinLocation(seeds []*seed) int {
	locationNums := []int{}
	for _, seed := range seeds {
		locationNums = append(locationNums, seed.locationNum)
	}
	return slices.Min(locationNums)
}

func calculateMinLocationPart2(seeds []*seedRange) int {
	locationNums := []int{}
	for _, seed := range seeds {
		locationNums = append(locationNums, seed.locationNum)
	}
	return slices.Min(locationNums)
}

func DayFivePartOne() {
	seeds, soilMap, fertilizerMap, waterMap, lightMap, tempMap, humdidityMap, locationMap := parseInput("./inputs/5_input.txt")
	for _, seed := range seeds {
		seed.setMapNums(
			soilMap,
			fertilizerMap,
			waterMap,
			lightMap,
			tempMap,
			humdidityMap,
			locationMap,
		)
		fmt.Println(seed)
	}
	fmt.Printf("Puzzle Output: %d\n", calculateMinLocation(seeds))
}

func DayFivePartTwo() {
	seeds, soilMap, fertilizerMap, waterMap, lightMap, tempMap, humdidityMap, locationMap := parseInputPart2("./inputs/5_input.txt")
	for _, seed := range seeds {
		seed.setMapNums(
			soilMap,
			fertilizerMap,
			waterMap,
			lightMap,
			tempMap,
			humdidityMap,
			locationMap,
		)
		fmt.Println(seed)
	}
	fmt.Printf("Puzzle Output: %d\n", calculateMinLocationPart2(seeds))
}
