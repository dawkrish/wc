package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("--------")

	var characterFlag bool
	var wordFlag bool
	var lineFlag bool
	var byteFlag bool

	flag.BoolVar(&characterFlag, "m", false, "Prints the number of characters")
	flag.BoolVar(&wordFlag, "w", false, "Prints the number of words")
	flag.BoolVar(&lineFlag, "l", false, "Prints the number of lines")
	flag.BoolVar(&byteFlag, "c", false, "Prints the number of bytes")

	flag.Parse()

	osArgs := flag.Args()

	for _, fileName := range osArgs {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("%v : No such file exists\n", fileName)
			continue
		}
		outputString := ""
		nWords := getWords(f)
		nBytes := getBytes(f)
		nLines, nChars := getLinesAndCharacters(fileName)

		if !characterFlag && !wordFlag && !lineFlag && !byteFlag {
			outputString += fmt.Sprintf("%d\t%d\t%d\t%v", nLines, nWords, nBytes, fileName)
		} else {
			if lineFlag {
				outputString += fmt.Sprintf("%d\t", nLines)
			}
			if wordFlag {
				outputString += fmt.Sprintf("%d\t", nWords)
			}
			if characterFlag {
				outputString += fmt.Sprintf("%d\t", nChars)
			}
			if byteFlag {
				outputString += fmt.Sprintf("%d\t", nBytes)
			}
			outputString += fmt.Sprintf("%v", fileName)

		}
		fmt.Println(outputString)
	}

	fmt.Println("--------")
}

func getWords(f *os.File) int {
	var w = 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w++
	}
	return w
}

func getLinesAndCharacters(fileName string) (int, int) {
	nC := 0
	bts, _ := os.ReadFile(fileName)
	content := string(bts)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		nC += len(line)
	}
	return len(strings.Split(content, "\n")) - 1, nC
}

func getBytes(f *os.File) int {
	fileInfo, _ := f.Stat()
	return int(fileInfo.Size())
}
