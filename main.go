package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
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

	for idx, fileName := range osArgs {
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
			if idx == 0 {
				fmt.Printf("L\tW\tB\tFILE\n")
			}
			outputString += fmt.Sprintf("%d\t%d\t%d\t%v", nLines, nWords, nBytes, fileName)
		} else {
			if lineFlag {
				if idx == 0 {
					fmt.Printf("L\t")
				}
				outputString += fmt.Sprintf("%d\t", nLines)
			}
			if wordFlag {
				if idx == 0 {
					fmt.Printf("W\t")
				}
				outputString += fmt.Sprintf("%d\t", nWords)
			}
			if characterFlag {
				if idx == 0 {
					fmt.Printf("C\t")
				}
				outputString += fmt.Sprintf("%d\t", nChars)
			}
			if byteFlag {
				if idx == 0 {
					fmt.Printf("B\t")
				}
				outputString += fmt.Sprintf("%d\t", nBytes)
			}
			if idx == 0 {
				fmt.Printf("FILE\n")
			}
			outputString += fmt.Sprintf("%v", fileName)

		}
		fmt.Println(outputString)
	}
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
		for range line {
			nC++
		}
		nC++ // because "\n" is also counted as a character
	}
	return len(lines) - 1, nC - 1 // because last "\n" is extra
}

func getBytes(f *os.File) int {
	fileInfo, _ := f.Stat()
	return int(fileInfo.Size())
}
