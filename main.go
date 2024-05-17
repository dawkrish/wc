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
		}
		nWords, nChars := getWordsAndCharacters(f)
		nBytes := getBytes(f)
		nLines := getLines(fileName)

	}

	fmt.Println("--------")
}

func getWordsAndCharacters(f *os.File) (int, int) {
	var w = 0
	var c = 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w++
		c += len(scanner.Text())
	}
	return w, c
}

func getLines(fileName string) int {
	bts, _ := os.ReadFile(fileName)
	content := string(bts)
	return len(strings.Split(content, "\n"))
}

func getBytes(f *os.File) int {
	fileInfo, _ := f.Stat()
	return int(fileInfo.Size())
}

func getCharacters() int {
	return 0
}
