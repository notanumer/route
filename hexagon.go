package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HexagonDrawer handles the drawing of ASCII hexagons
type HexagonDrawer struct {
	width  int
	height int
}

// NewHexagonDrawer creates a new hexagon drawer with given dimensions
func NewHexagonDrawer(width, height int) *HexagonDrawer {
	return &HexagonDrawer{
		width:  width,
		height: height,
	}
}

// Draw generates the hexagon lines as a slice of strings
func (h *HexagonDrawer) Draw() []string {
	var lines []string
	
	// Top part of hexagon (expanding outward)
	for i := 0; i < h.height; i++ {
		spacesBefore := h.height - 1 - i
		spacesBetween := h.width + 2*i
		line := strings.Repeat(" ", spacesBefore) + "/" + 
				strings.Repeat(" ", spacesBetween) + "\\"
		lines = append(lines, line)
	}
	
	// Middle part (widest part with underscores)
	middleWidth := h.width + 2*h.height
	middleLine := "/" + strings.Repeat("_", middleWidth) + "\\"
	lines = append(lines, middleLine)
	
	// Bottom part of hexagon (contracting inward)
	for i := 0; i < h.height; i++ {
		spacesBefore := i
		underscoresBetween := h.width + 2*(h.height-1-i)
		line := strings.Repeat(" ", spacesBefore) + "\\" + 
				strings.Repeat("_", underscoresBetween) + "/"
		lines = append(lines, line)
	}
	
	return lines
}

// parseInput reads and parses the input
func parseInput() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Read number of test cases
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read number of test cases")
	}
	
	t, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return nil, fmt.Errorf("invalid number of test cases: %v", err)
	}
	
	var testCases [][]int
	
	// Read each test case
	for i := 0; i < t; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("failed to read test case %d", i+1)
		}
		
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid input format for test case %d", i+1)
		}
		
		width, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid width in test case %d: %v", i+1, err)
		}
		
		height, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid height in test case %d: %v", i+1, err)
		}
		
		testCases = append(testCases, []int{width, height})
	}
	
	return testCases, nil
}

func main() {
	testCases, err := parseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Process each test case
	for i, testCase := range testCases {
		width, height := testCase[0], testCase[1]
		
		// Add empty line between hexagons (except before the first one)
		if i > 0 {
			fmt.Println()
		}
		
		// Create and draw hexagon
		drawer := NewHexagonDrawer(width, height)
		lines := drawer.Draw()
		
		// Print all lines of the hexagon
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}