package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	lineNumer := 1
	p := make(map[int]struct{})
	p1 := make(map[int]struct{})
	for fileScanner.Scan() {
		textLine := fileScanner.Text()
		for i := range textLine {
			priority := int(textLine[i])
			if priority >= 65 && priority < 91 {
				// Capital letter
				priority = priority - 65 + 27
			} else if priority >= 97 && priority < 123 {
				// Lower letter
				priority = priority - 97 + 1
			}

			if lineNumer%3 == 0 {
				if _, ok := p1[priority]; ok {
					sum += priority
					// reset the maps for the next round of 3 lines
					p = make(map[int]struct{})
					p1 = make(map[int]struct{})
					break
				}
			} else if lineNumer%3 == 2 {
				if _, ok := p[priority]; ok {
					// Elements common between line 1 and 2
					p1[priority] = struct{}{}
				}
			} else if lineNumer%3 == 1 {
				p[priority] = struct{}{}
			}
		}
		lineNumer += 1
	}

	fmt.Println(sum)
}
