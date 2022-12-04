package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)
	i := 0
	for fileScanner.Scan() {
		textLine := string(fileScanner.Text())
		bothRanges := strings.Split(textLine, ",")
		range1 := strings.Split(bothRanges[0], "-")
		range2 := strings.Split(bothRanges[1], "-")

		range10, _ := strconv.Atoi(range1[0])
		range20, _ := strconv.Atoi(range2[0])
		range11, _ := strconv.Atoi(range1[1])
		range21, _ := strconv.Atoi(range2[1])

		if (range10 >= range20 && range11 <= range21) ||
			(range20 >= range10 && range21 <= range11) {
			i++
		}
	}

	fmt.Println(i)
}
