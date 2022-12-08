package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rootFile = "/"

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	currentFile := rootFile

	folders := make(map[string]int)
	parentFolders := make(map[string]string)

	for fileScanner.Scan() {
		textLine := fileScanner.Text()
		var actualFile string
		fmt.Sscanf(textLine, "$ cd %s", &actualFile)
		if actualFile != "" {
			if actualFile == ".." {
				currentFile = parentFolders[currentFile]
			} else if actualFile == rootFile {
				currentFile = rootFile
			} else {
				link := "/"
				if currentFile == rootFile {
					link = ""
				}
				currentFile = currentFile + link + actualFile
			}
			continue
		} else {
			if string(textLine) == "$ ls" {
				continue
			}
			words := strings.Split(string(textLine), " ")
			if words[0] == "dir" {
				link := "/"
				if currentFile == rootFile {
					link = ""
				}
				folderPath := currentFile + link + words[1]
				parentFolders[folderPath] = currentFile
				continue
			}

			size, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}

			fileToAddSize := currentFile
			for true {
				parent, isParentExit := parentFolders[fileToAddSize]
				if !isParentExit {
					break
				}
				folders[fileToAddSize] += size
				fileToAddSize = parent
			}
		}
	}

	total := 0
	for key, folderSize := range folders {
		if strings.Count(key, "/") == 1 {
			total += folderSize
		}
	}

	totalSpace := 70000000
	unusedSpaceNeeded := 30000000
	currentUnusedSpace := totalSpace - total
	spaceToDelete := unusedSpaceNeeded - currentUnusedSpace
	if spaceToDelete <= 0 {
		fmt.Println(0)
		return
	}

	min := total
	for _, folderSize := range folders {
		if folderSize >= spaceToDelete && folderSize <= min {
			min = folderSize
		}
	}

	fmt.Println(min)
}
