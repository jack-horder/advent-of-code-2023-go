package d8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	label    string
	leftStr  string
	rightStr string
	left     *node
	right    *node
}

func (n node) String() string {
	return fmt.Sprintf("<label: %v, left: %v, right: %v>", n.label, n.left.label, n.right.label)
}

type nodeMap map[string]*node

func parseNodes(lines []string) []*node {
	nodes := []*node{}
	for _, line := range lines {
		n := new(node)
		lineSplit := strings.Split(line, " = ")
		label := lineSplit[0]
		leftRightNodes := strings.ReplaceAll(lineSplit[1], "(", "")
		leftRightNodes = strings.ReplaceAll(leftRightNodes, ")", "")
		nodeSlice := strings.Split(leftRightNodes, ", ")
		n.label = label
		n.leftStr = nodeSlice[0]
		n.rightStr = nodeSlice[1]
		nodes = append(nodes, n)
	}
	return nodes
}

func parseNodeStrings(nodes []*node, nm nodeMap) {
	for _, n := range nodes {
		n.left = nm[n.leftStr]
		n.right = nm[n.rightStr]
	}
}

func createNodeMap(nodes []*node) map[string]*node {
	nodeMap := make(map[string]*node)
	for _, n := range nodes {
		nodeMap[n.label] = n
	}
	return nodeMap
}

func traverseNode(n *node, instruction rune) *node {
	switch string(instruction) {
	case "L":
		return n.left
	case "R":
		return n.right
	default:
		return nil
	}
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

func getStartNodes(nodes []*node) []*node {
	var startNodes []*node
	for _, n := range nodes {
		if strings.HasSuffix(n.label, "A") {
			startNodes = append(startNodes, n)
		}
	}
	return startNodes
}

func checkEndState(n *node) bool {
	return strings.HasSuffix(n.label, "Z")
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func DayEightPartOne() {
	lines, err := readLines("./inputs/8_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	instructions := []rune(lines[0])
	nodes := parseNodes(lines[2:])
	nm := createNodeMap(nodes)
	parseNodeStrings(nodes, nm)

	currentNode := nm["AAA"]
	steps := 0
	for currentNode.label != "ZZZ" {
		for _, i := range instructions {
			currentNode = traverseNode(currentNode, i)
			steps++
		}
	}
	fmt.Printf("Puzzle Output: %d\n", steps)
}

func DayEightPartTwo() {
	lines, err := readLines("./inputs/8_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	instructions := []rune(lines[0])
	nodes := parseNodes(lines[2:])
	nm := createNodeMap(nodes)
	parseNodeStrings(nodes, nm)
	startNodes := getStartNodes(nodes)
	currentIterations := make([]int, len(startNodes))
	for idx, n := range startNodes {
		endNode := n
		for !checkEndState(endNode) {
			for _, i := range instructions {
				endNode = traverseNode(endNode, i)
			}
			currentIterations[idx] += 1
		}
	}
	for i := 0; i < len(currentIterations); i++ {
		currentIterations[i] = currentIterations[i] * len(instructions)
	}
	// for !checkEndState(currentNodes) {
	// 	for _, i := range instructions {
	// 		currentNodes = traverseNodes(currentNodes, i)
	// 		steps++
	// 		fmt.Println(currentNodes)
	// 		if checkEndState(currentNodes) {
	// 			break
	// 		}
	// 	}
	// }
	fmt.Printf("Puzzle Output: %d\n", LCM(currentIterations[0], currentIterations[1], currentIterations[2:]...))
}
