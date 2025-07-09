package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checksum(runes []rune) int {
	res := 0
	for idx, char := range runes {
		if char == '.' {
			continue
		}
		num := int(char) - '0' // convert rune to int and then ascii remove until we get nums 0-9
		res += num * idx
	}
	return res
}

// swaps 2 nums and returns the stringbuilder swapped
func swap(i int, j int, runes []rune) []rune {
	runes[i], runes[j] = runes[j], runes[i]

	return runes
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)
	var lineStr string
	var sb strings.Builder

	// input is 1 line
	if scanner.Scan() {
		lineStr = scanner.Text()
	}

	id := 0

	for idx, char := range lineStr {
		if idx%2 == 0 {
			// file
			count, err := strconv.Atoi(string(char))
			check(err)
			for range count {
				_, err := sb.WriteString(strconv.Itoa(id))
				check(err)
			}
			id++
		} else {
			// add '.'
			count, err := strconv.Atoi(string(char))
			check(err)
			for range count {
				_, err := sb.WriteRune('.')
				check(err)
			}
		}
	}

	runes := []rune(sb.String())
	fmt.Println("Initial disk state:", string(runes))

	// begin 2 pointer
	head := 0
	tail := sb.Len() - 1

	for head < tail {
		if runes[head] == '.' && runes[tail] != '.' {
			runes = swap(head, tail, runes)
			head++
			tail--
		} else if runes[head] != '.' {
			head++
		} else if runes[tail] == '.' {
			tail--
		}
	}

	fmt.Println("Final disk state:", string(runes))
	fmt.Println("Checksum:", checksum(runes))

}
