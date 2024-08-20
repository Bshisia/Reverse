package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]
	banner := ""
	if len(os.Args) == 2 {
		banner = "standard.txt"
	} else if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	if !strings.Contains(arg, "--reverse=") {
		return
	}

	bannerFile, err := ReadFile(banner)
	if err {
		return
	}
	slicedBanner := SliceFile(bannerFile)
	file, err := ReadFile(arg[10:])
	if err {
		return
	}

	str := ""
	for len(file[0]) > 0 {
		for i, val := range slicedBanner {
			if CheckPattern(val, file) {
				str += string(rune(i + 32))
				file = TrimFound(len(val[0]), file)
			}
		}
	}
	fmt.Println(str)
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
	if len(char[0]) > len(word[0]) {
		return false
	}
	for i, str := range word[:len(word)-1] {
		if char[i] != str[:len(char[i])] {
			return false
		}
	}
	return true
}

func TrimFound(length int, word []string) []string {
	for i, val := range word[0 : len(word)-1] {
		// fmt.Println(val[length:])
		word[i] = val[length:]
	}
	return word
}
