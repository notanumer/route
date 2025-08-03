package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var width, height int
		fmt.Fscan(in, &width, &height)

		// Top line
		topSpaces := height - 1
		topLine := strings.Repeat(" ", topSpaces) + strings.Repeat("_", width)
		fmt.Fprintln(out, topLine)

		// Upper diagonal lines
		for j := 0; j < height-1; j++ {
			spacesBefore := height - 1 - j - 1
			spacesBetween := width + 2*j
			line := strings.Repeat(" ", spacesBefore) + "/" + strings.Repeat(" ", spacesBetween) + "\\"
			fmt.Fprintln(out, line)
		}

		// Middle line
		middleLine := "/" + strings.Repeat("_", width + 2*(height-1)) + "\\"
		fmt.Fprintln(out, middleLine)

		// Lower diagonal lines
		for j := 0; j < height-1; j++ {
			spacesBefore := j + 1
			spacesBetween := width + 2*(height-2-j)
			line := strings.Repeat(" ", spacesBefore) + "\\" + strings.Repeat(" ", spacesBetween) + "/"
			fmt.Fprintln(out, line)
		}

		// Bottom line
		bottomSpaces := height - 1
		bottomLine := strings.Repeat(" ", bottomSpaces) + strings.Repeat("_", width)
		fmt.Fprintln(out, bottomLine)

		// Add an empty line between hexagons except the last one
		if i < t-1 {
			fmt.Fprintln(out)
		}
	}
}