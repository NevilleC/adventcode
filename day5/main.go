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

	stacks := make([][]string, 9)

	buildStack := true
	for fileScanner.Scan() {
		textLine := string(fileScanner.Text())
		i := 0
		for i < 9 && buildStack {
			el := string(textLine[1+i*4])
			if el == "1" {
				buildStack = false
				fileScanner.Scan() // to skip the empty line
				break
			}

			if textLine[1+i*4] >= 65 && textLine[1+i*4] < 91 {
				// To ensure the new element is getting before
				// the previous ones
				stacks[i] = append([]string{el}, stacks[i]...)
			}
			i++
		}

		if !buildStack {
			var move, from, to int
			fmt.Sscanf(textLine, "move %d from %d to %d", &move, &from, &to)
			j := 0
			for j < move && move > 0 {
				stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
				stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
				j++
			}
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
}
