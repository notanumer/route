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
	
	// Simple sequential placement to get exactly k hexagons
	for row := 0; placed < k; row++ {
		y := row * (height + 1)
		if y + 2*height + 1 > n {
			break
		}
		
		// Hexagonal grid offset
		offset := 0
		if row%2 == 1 {
			offset = (width + height) / 2
		}
		
		for col := 0; placed < k; col++ {
			x := col * (width + height) + offset
			if x + width + 2*height > m {
				break
			}
			
			drawHexagon(screen, x+1, y+1, width, height, m, n)
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

func drawHexagon(screen [][]rune, startX, startY, width, height, maxX, maxY int) {
	// Adjust coordinates for screen border
	startX++
	startY++
	
	// Top line of underscores
	for i := 0; i < width; i++ {
		x := startX + height + i
		y := startY
		if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
			screen[y][x] = '_'
		}
	}
	
	// Upper expanding sides
	for i := 0; i < height; i++ {
		y := startY + 1 + i
		
		// Left side: /
		x := startX + height - 1 - i
		if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
			screen[y][x] = '/'
		}
		
		// Right side: \
		x = startX + height + width + i
		if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
			screen[y][x] = '\\'
		}
	}
	
	// Lower contracting sides
	for i := 0; i < height; i++ {
		y := startY + height + 1 + i
		
		// Left side: \
		x := startX + i
		if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
			screen[y][x] = '\\'
		}
		
		// Right side: /
		x = startX + width + 2*height - 1 - i
		if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
			screen[y][x] = '/'
		}
		
		// Bottom line on last iteration
		if i == height-1 {
			for j := 0; j < width; j++ {
				x = startX + height + j
				if x >= 1 && x <= maxX && y >= 1 && y <= maxY {
					screen[y][x] = '_'
				}
			}
		}
	}
}