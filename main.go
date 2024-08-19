package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	arg := os.Args[1:]
	if !strings.Contains(arg[0], "--reverse=") {
		return
	}
	bannerFile, err := ReadFile(arg[1])
	if err {
		return
	}
	slicedBanner := SliceFile(bannerFile)
	file, err := ReadFile(arg[0][10:])
	if err {
		return
	}
	arg := ""
	for len(file) > 0 {
		for i, val := range slicedBanner {
			if CheckPattern(val, file) {
				str += string(rune(i + 32))
			}
		}
	}
}

func ReadFile(fileName string) ([]string, bool) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error %d\n", err)
		return nil, true
	}
	lines := strings.Split(string(content), "\n")
	return lines, false
}

func SliceFile(fileSlice []string) [][]string {
	var result [][]string
	for i := 0; i < len(fileSlice)-1; i += 9 {
		file := fileSlice[i+1 : i+9]
		result = append(result, file)
	}
	return result
}

func CheckPattern(char, word []string) bool {
	present := true
	if len(char[0]) > len(word) {
		return false
	}
	for i, str := range word[0 : len(word)-1] {
		if char[i] != str[:len(char[i])] {
			present = false
		}
	}
	return present
}
