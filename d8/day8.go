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
