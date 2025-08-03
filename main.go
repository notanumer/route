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

	// Initialize screen with borders
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
	
	// Правильная сотовая структура на основе анализа ожидаемого результата:
	// Вертикальные шестиугольники должны разделять горизонтальные линии
	rowSpacing := 2 * height
	colSpacing := 2 * (width + height)

	// Store positions of hexagons for edge sharing - removed for simplicity
	
	for row := 0; placed < k; row++ {
		startY := row * rowSpacing
		
		// Check if hexagon fits vertically
		if startY + 2*height + 1 > n {
			break
		}

		// Honeycomb offset for odd rows
		offsetX := 0
		if row%2 == 1 {
			// Применяем смещение только если помещается больше одного шестиугольника
			maxHexagonsInRow := m / colSpacing
			if maxHexagonsInRow > 1 {
				offsetX = colSpacing / 2
			}
		}

		for col := 0; placed < k; col++ {
			startX := col * colSpacing + offsetX
			
			// Check if hexagon fits horizontally
			if startX + width + 2*height > m {
				break
			}

			// Draw hexagon accounting for border
			// Position hexagon so that first hexagon has elements in first row/column
			hexX := startX + 1
			hexY := startY + 1
			
			drawHexagon(screen, hexX, hexY, width, height)
			placed++
		}
	}

	// Соединяем соседние шестиугольники
	connectHexagons(screen, m, n)

	// Print the screen
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Print(string(screen[i][j]))
		}
		fmt.Println()
	}
}

func drawHexagon(screen [][]rune, x, y, width, height int) {
	// Top edge 
	for i := 0; i < width; i++ {
		screen[y][x+height+i] = '_'
	}

	// Upper slanted edges
	for i := 0; i < height; i++ {
		rowIdx := y + 1 + i
		screen[rowIdx][x+height-1-i] = '/'
		screen[rowIdx][x+height+width+i] = '\\'
	}

	// Lower slanted edges
	for i := 0; i < height; i++ {
		rowIdx := y + height + 1 + i
		screen[rowIdx][x+i] = '\\'
		screen[rowIdx][x+width+2*height-1-i] = '/'
	}

	// Bottom edge
	for i := 0; i < width; i++ {
		screen[y+2*height][x+height+i] = '_'
	}
}

// Функция для соединения соседних шестиугольников
func connectHexagons(screen [][]rune, m, n int) {
	// Проходим по экрану и ищем места, где нужно соединить диагональные рёбра
	for i := 1; i <= n; i++ {
		for j := 1; j <= m-2; j++ {
			// Ищем узор "\ /" и заменяем на "\_/" только если это правильное соединение
			if screen[i][j] == '\\' && screen[i][j+1] == ' ' && screen[i][j+2] == '/' {
				screen[i][j+1] = '_'
			}
		}
	}
}