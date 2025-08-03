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
	
	// Place hexagons in a pattern that matches the expected output
	placed := 0
	
	// For the pattern, each row has hexagons spaced by (width + height)
	// Odd rows are offset
	rowHeight := height + 1
	colWidth := width + height
	
	for row := 0; placed < k; row++ {
		y := row * rowHeight
		if y + 2*height + 1 > n {
			break
		}
		
		// Calculate starting x position for this row
		startX := 0
		if row%2 == 1 {
			startX = colWidth / 2
		}
		
		for col := 0; placed < k; col++ {
			x := startX + col * colWidth
			
			// Check bounds
			if x + width + 2*height > m {
				break
			}
			
			// Draw single hexagon
			if drawSingleHex(screen, x+1, y+1, width, height, m, n) {
				placed++
			} else {
				break
			}
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

func drawSingleHex(screen [][]rune, startX, startY, width, height, maxX, maxY int) bool {
	// Check if hexagon fits completely within bounds
	if startX+width+2*height > maxX+1 || startY+2*height+1 > maxY+1 {
		return false
	}
	
	// Top line
	for i := 0; i < width; i++ {
		if startX+height+i <= maxX {
			screen[startY][startX+height+i] = '_'
		}
	}
	
	// Upper part
	for i := 0; i < height; i++ {
		y := startY + 1 + i
		if y <= maxY {
			if startX+height-1-i >= 1 {
				screen[y][startX+height-1-i] = '/'
			}
			if startX+height+width+i <= maxX {
				screen[y][startX+height+width+i] = '\\'
			}
		}
	}
	
	// Lower part
	for i := 0; i < height; i++ {
		y := startY + height + 1 + i
		if y <= maxY {
			if startX+i >= 1 {
				screen[y][startX+i] = '\\'
			}
			if startX+width+2*height-1-i <= maxX {
				screen[y][startX+width+2*height-1-i] = '/'
			}
			
			// Bottom line
			if i == height-1 {
				for j := 0; j < width; j++ {
					if startX+height+j <= maxX {
						screen[y][startX+height+j] = '_'
					}
				}
			}
		}
	}
	
	return true
}