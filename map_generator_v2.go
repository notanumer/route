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
	hexWidth := width + 2*height
	
	// Hexagonal tiling: adjacent hexagons share sides
	// Horizontal step = width + height (sharing sides)
	// Vertical step = height + 1
	stepX := width + height
	stepY := height + 1
	
	for row := 0; placed < k; row++ {
		y := row * stepY
		if y + 2*height + 1 > n {
			break
		}
		
		// Offset for odd rows (hexagonal tiling pattern)
		offset := 0
		if row%2 == 1 {
			offset = stepX / 2
		}
		
		for col := 0; placed < k; col++ {
			x := col * stepX + offset
			if x + hexWidth > m {
				break
			}
			
			drawHexagon(screen, x+1, y+1, width, height)
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

func drawHexagon(screen [][]rune, startX, startY, width, height int) {
	// Top line
	for i := 0; i < width; i++ {
		screen[startY][startX+height+i] = '_'
	}
	
	// Upper part
	for i := 0; i < height; i++ {
		y := startY + 1 + i
		screen[y][startX+height-1-i] = '/'
		screen[y][startX+height+width+i] = '\\'
	}
	
	// Lower part
	for i := 0; i < height; i++ {
		y := startY + height + 1 + i
		screen[y][startX+i] = '\\'
		screen[y][startX+width+2*height-1-i] = '/'
		
		// Bottom line
		if i == height-1 {
			for j := 0; j < width; j++ {
				screen[y][startX+height+j] = '_'
			}
		}
	}
}