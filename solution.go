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
	rowSpacing := height + 1
	if height == 1 {
		rowSpacing = 2  // For height=1, use tighter spacing
	}
	colSpacing := width + height
	
	// Place hexagons row by row with hexagonal tiling
	// Shift the entire grid so that the first hexagon starts at (0,0)
	shiftX := 1 - height  // Move left so the leftmost part of first hexagon is at x=1
	shiftY := 1           // Start at the top
	
	for row := 0; placed < k; row++ {
		y := row * rowSpacing + shiftY
		if y + 2*height + 1 > n+1 {
			break
		}
		
		// Offset for odd rows (hexagonal tiling)
		offset := 0
		if row%2 == 1 {
			offset = colSpacing / 2
		}
		
		for col := 0; placed < k; col++ {
			x := col * colSpacing + offset + shiftX
			
			// Check if hexagon would fit reasonably within bounds
			if x + width + height > m + height {
				break
			}
			
			// Draw hexagon
			drawHexagon(screen, x, y, width, height)
			placed++
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

func drawHexagon(screen [][]rune, startX, startY, width, height int) {
	// Draw hexagon starting at (startX, startY)
	// The first hexagon should have some symbols at position (0,0)
	
	// Top line: underscores
	for i := 0; i < width; i++ {
		if startX+height+i >= 1 && startX+height+i < len(screen[0])-1 {
			screen[startY][startX+height+i] = '_'
		}
	}
	
	// Upper expanding part
	for i := 0; i < height; i++ {
		y := startY + 1 + i
		if y >= 1 && y < len(screen)-1 {
			// Left side /
			x := startX + height - 1 - i
			if x >= 1 && x < len(screen[0])-1 {
				screen[y][x] = '/'
			}
			// Right side \
			x = startX + height + width + i
			if x >= 1 && x < len(screen[0])-1 {
				screen[y][x] = '\\'
			}
		}
	}
	
	// Lower contracting part
	for i := 0; i < height; i++ {
		y := startY + height + 1 + i
		if y >= 1 && y < len(screen)-1 {
			// Left side \
			x := startX + i
			if x >= 1 && x < len(screen[0])-1 {
				screen[y][x] = '\\'
			}
			// Right side /
			x = startX + width + 2*height - 1 - i
			if x >= 1 && x < len(screen[0])-1 {
				screen[y][x] = '/'
			}
			
			// Bottom line: underscores on last iteration
			if i == height-1 {
				for j := 0; j < width; j++ {
					x = startX + height + j
					if x >= 1 && x < len(screen[0])-1 {
						screen[y][x] = '_'
					}
				}
			}
		}
	}
}