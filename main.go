package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	parts := strings.Fields(scanner.Text())
	if len(parts) != 5 {
		return
	}

	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])
	_, _ = strconv.Atoi(parts[2])  // width
	height, _ := strconv.Atoi(parts[3])
	k, _ := strconv.Atoi(parts[4])

	// Create screen with borders
	screen := make([][]rune, n+2)
	for i := range screen {
		screen[i] = make([]rune, m+2)
		for j := range screen[i] {
			if i == 0 || i == n+1 {
				if j == 0 || j == m+1 {
					screen[i][j] = '+'
				} else {
					screen[i][j] = '-'
				}
			} else {
				if j == 0 || j == m+1 {
					screen[i][j] = '|'
				} else {
					screen[i][j] = ' '
				}
			}
		}
	}

	placed := 0
	
	// Place hexagons in honeycomb pattern
	for row := 0; placed < k; row++ {
		y := 1 + row * (height + 1)
		if y + 2*height >= n+1 {
			break
		}
		
		// Calculate horizontal offset for honeycomb pattern
		offsetX := 0
		if row%2 == 1 {
			offsetX = 2  // From analysis: odd rows start at position 3 (offset +2)
		}
		
		hexagonsInRow := 0
		// Count hexagons that fit in this row
		for col := 0; placed + hexagonsInRow < k; col++ {
			x := 1 + col * 4 + offsetX
			// Check if the full hexagon pattern fits (needs 3 characters: \_/)
			if x + 2 >= m+1 {
				break
			}
			hexagonsInRow++
		}
		
		// Draw the row pattern
		drawHoneycombRow(screen, y, hexagonsInRow, offsetX)
		placed += hexagonsInRow
	}

	// Print the screen
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Print(string(screen[i][j]))
		}
		fmt.Println()
	}
}

func drawHoneycombRow(screen [][]rune, startY, count, offsetX int) {
	// Row 0: Top edges at positions 1,5,9,13,17
	for i := 0; i < count; i++ {
		x := 1 + i*4 + offsetX
		screen[startY][x+1] = '_'
	}
	
	// Row 1: Pattern |/ \_/ \_/ \_/ \_/ \ |
	if count > 0 {
		// First hexagon: / at pos 1, \ at pos 3
		x := 1 + offsetX
		screen[startY+1][x] = '/'
		screen[startY+1][x+2] = '\\'
		
		// Subsequent patterns: _/ \
		for i := 1; i < count; i++ {
			x = 1 + i*4 + offsetX
			screen[startY+1][x-1] = '_'
			screen[startY+1][x] = '/'    
			screen[startY+1][x+2] = '\\'
		}
	}
	
	// Row 2: Pattern |\_/ \_/ \_/ \_/ \_/ |
	for i := 0; i < count; i++ {
		x := 1 + i*4 + offsetX
		screen[startY+2][x] = '\\'
		screen[startY+2][x+1] = '_'
		screen[startY+2][x+2] = '/'
		// Clear the space position that might have been overwritten
		if i < count-1 {
			screen[startY+2][x+3] = ' '
		}
	}
}