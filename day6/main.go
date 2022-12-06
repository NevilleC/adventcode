package main

import (
	"bufio"
	"fmt"
	"os"
)

const diffChars = 14

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	textLine := fileScanner.Text()
	for i := range textLine {
		seq := make(map[string]struct{})
		for j := 0; j < diffChars; j++ {
			seq[string(textLine[i+j])] = struct{}{}
		}

		if len(seq) == diffChars {
			fmt.Println(i + diffChars)
			break
		}
	}
}
