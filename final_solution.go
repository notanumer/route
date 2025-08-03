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
	
	// Create screen with border
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
	
	// For hexagonal tiling with shared sides
	colStep := width + height  // Distance between hexagon centers horizontally
	rowStep := height + 1      // Distance between rows vertically
	
	for row := 0; placed < k; row++ {
		y := row * rowStep + 1
		
		// Check if row fits within screen
		if y + 2*height + 1 > n+1 {
			break
		}
		
		// Horizontal offset for odd rows (hexagonal tiling)
		offsetX := 0
		if row%2 == 1 {
			offsetX = colStep / 2
		}
		
		for col := 0; placed < k; col++ {
			x := col * colStep + offsetX + 1
			
			// Check if hexagon fits horizontally
			if x + width + 2*height - 1 > m+1 {
				break
			}
			
			// Draw hexagon at position (x, y)
			drawHex(screen, x, y, width, height)
			placed++
		}
	}
	
	// Print screen
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Print(string(screen[i][j]))
		}
		fmt.Println()
	}
}

func drawHex(screen [][]rune, startX, startY, width, height int) {
	maxY := len(screen) - 1
	maxX := len(screen[0]) - 1
	
	// Top line with underscores
	for i := 0; i < width; i++ {
		x := startX + height + i
		if x >= 1 && x <= maxX {
			screen[startY][x] = '_'
		}
	}
	
	// Upper expanding part
	for i := 0; i < height; i++ {
		y := startY + 1 + i
		if y >= 1 && y <= maxY {
			// Left side /
			x := startX + height - 1 - i
			if x >= 1 && x <= maxX {
				screen[y][x] = '/'
			}
			// Right side \
			x = startX + height + width + i
			if x >= 1 && x <= maxX {
				screen[y][x] = '\\'
			}
		}
	}
	
	// Lower contracting part
	for i := 0; i < height; i++ {
		y := startY + height + 1 + i
		if y >= 1 && y <= maxY {
			// Left side \
			x := startX + i
			if x >= 1 && x <= maxX {
				screen[y][x] = '\\'
			}
			// Right side /
			x = startX + width + 2*height - 1 - i
			if x >= 1 && x <= maxX {
				screen[y][x] = '/'
			}
			
			// Bottom line on last iteration
			if i == height-1 {
				for j := 0; j < width; j++ {
					x = startX + height + j
					if x >= 1 && x <= maxX {
						screen[y][x] = '_'
					}
				}
			}
		}
	}
}