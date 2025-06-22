package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)

	var sum int
	const grammar_1 = `mul\(\d{1,3},\d{1,3}\)|don\'t\(\)|do\(\)`
	const grammar_2 = `(\d{1,3}),(\d{1,3})`

	r1, err := regexp.Compile(grammar_1)
	check(err)
	r2, err := regexp.Compile(grammar_2)
	check(err)

	var do bool = true // starts as do()

	for scanner.Scan() {
		corrupted := scanner.Text()

		muls := r1.FindAllString(corrupted, -1) // part1 format: ["mul(843,597)", "mul(717, 524)", ...]
		// part2 format: ["mul(843, 597)", "do()", ...]
		for _, mul := range muls {
			switch mul {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if !do {
					continue
				}

				numsStr := r2.FindString(mul)          // e.g., "843,597"
				numsArr := strings.Split(numsStr, ",") // ["843", "597"]

				num1, err := strconv.Atoi(numsArr[0])
				check(err)
				num2, err := strconv.Atoi(numsArr[1])
				check(err)

				sum = sum + num1*num2
			}
		}
	}

	fmt.Println(sum)
}
