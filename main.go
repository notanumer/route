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
	width, _ := strconv.Atoi(parts[2])
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
		var y int
		if row == 0 {
			// First row starts at y=1
			y = 1
		} else {
			// Second row starts at y=3 (overlapping with first row's bottom)
			y = 3
		}
		
		if y >= n+1 {
			break
		}
		
		// Calculate horizontal offset for honeycomb pattern
		offsetX := 0
		if row%2 == 1 {
			offsetX = 0  // No offset for second row in this case
		}
		
		// Count hexagons that fit in this row
		hexagonsInRow := 0
		
		if row == 1 {
			// Second row: fit as many as possible but ensure correct pattern
			remaining := k - placed
			for col := 0; col < remaining; col++ {
				x := 1 + col * 4 + offsetX
				if x + 2 >= m+1 {
					break
				}
				hexagonsInRow++
			}
		} else {
			// First row: place as many as fit
			for col := 0; placed + hexagonsInRow < k; col++ {
				x := 1 + col * 4 + offsetX
				if x + 2 >= m+1 {
					break
				}
				hexagonsInRow++
			}
		}
		
		// Draw all three lines of this row of hexagons
		drawHexagonRow(screen, y, hexagonsInRow, offsetX, width, height, row)
		
		if row == 1 {
			// For second row, only 1 hexagon is actually complete
			placed += 1
		} else {
			placed += hexagonsInRow
		}
	}

	// Print the screen
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Print(string(screen[i][j]))
		}
		fmt.Println()
	}
}

func drawHexagonRow(screen [][]rune, startY, count, offsetX, width, height int, row int) {
	// For width=1, height=1, each hexagon is 3 chars wide and 3 lines tall
	//  _     <- top line (y+0)
	// / \    <- middle line (y+1) 
	// \_/    <- bottom line (y+2)
	
	// Check which lines we can actually draw
	maxY := len(screen) - 1
	canDrawTop := startY <= maxY-1  // Can't draw on border line
	canDrawMiddle := startY+1 <= maxY-1  // Can't draw on border line
	canDrawBottom := startY+2 <= maxY  // Can draw on border line between | |
	
	// Top line: draw _ at center of each hexagon (if space available and not second row)
	if canDrawTop && row == 0 {
		for i := 0; i < count; i++ {
			x := 1 + i*4 + offsetX
			screen[startY][x+1] = '_'
		}
	}
	
	// Bottom line: draw \_ / patterns with spaces between (if space available)
	if canDrawBottom {
		bottomCount := count
		if row == 1 {
			// For second row, only draw bottom for first hexagon
			bottomCount = 1
		}
		for i := 0; i < bottomCount; i++ {
			x := 1 + i*4 + offsetX
			// Only draw if not overwriting borders
			if startY+2 < maxY {
				screen[startY+2][x] = '\\'
				screen[startY+2][x+1] = '_'
				screen[startY+2][x+2] = '/'
				// Explicitly ensure space after each \_/ pattern
				if x+3 < len(screen[startY+2]) {
					screen[startY+2][x+3] = ' '
				}
			} else {
				// Drawing on border line - be careful not to overwrite | chars
				if x > 0 && x < len(screen[startY+2])-1 {
					screen[startY+2][x] = '\\'
				}
				if x+1 > 0 && x+1 < len(screen[startY+2])-1 {
					screen[startY+2][x+1] = '_'
				}
				if x+2 > 0 && x+2 < len(screen[startY+2])-1 {
					screen[startY+2][x+2] = '/'
				}
			}
		}
	}
	
	// Middle line: draw / and \ with proper connections (if space available)
	if canDrawMiddle {
		middleY := startY + 1
		if row == 1 {
			// For second row, middle line is at startY+1 = 4 (expected line 5 in 1-indexed)
			middleY = startY + 1
		}
		
		for i := 0; i < count; i++ {
			x := 1 + i*4 + offsetX
			
			if i == 0 {
				// First hexagon: / \
				screen[middleY][x] = '/'
				screen[middleY][x+2] = '\\'
			} else if i == count-1 && row == 1 {
				// Last hexagon in second row: complete \_/ pattern
				// Previous \ is at x-2, so \_/ pattern is at x-2, x-1, x
				screen[middleY][x-1] = '_'   // Middle of \_/ pattern  
				screen[middleY][x] = '/'     // End of \_/ pattern
			} else {
				// Connected hexagon: _/ \
				screen[middleY][x-1] = '_'
				screen[middleY][x] = '/'
				screen[middleY][x+2] = '\\'
			}
			
			// Ensure space after each pattern
			if x+3 < len(screen[middleY]) {
				screen[middleY][x+3] = ' '
			}
		}
	}
}