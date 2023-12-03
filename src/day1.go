package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	readFile, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	p1 := []int{}
	p2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		p1 = append(p1, partOne(line))
		p2 = append(p2, partTwo(line))
	}

	p1t := 0
	for _, s := range p1 {
		p1t = p1t + s
	}

	p2t := 0
	for _, s := range p2 {
		p2t = p2t + s
	}
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", p1t, p2t)

}

func partTwo(line string) int {
	return appendNums(findFirst(line), findLast(line))
}

func findFirst(line string) int {
	for i := 0; i < len(line); i++ {
		if num, err := strconv.Atoi(string(line[i])); err == nil {
			return num
		}
		part := line[0 : i+1]
		for i, n := range numbers {
			if strings.Contains(part, n) {
				return i + 1
			}
		}
	}
	log.Fatal("didn't find a number: ", line)
	return 0
}

func findLast(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if num, err := strconv.Atoi(string(line[i])); err == nil {
			return num
		}
		part := line[i:]
		for x, n := range numbers {
			if strings.Contains(part, n) {
				return x + 1
			}
		}
	}
	log.Fatal("didn't find a number: ", line)
	return 0
}

func partOne(line string) int {
	digits := []int{}
	for i := 0; i < len(line); i++ {
		if num, err := strconv.Atoi(string(line[i])); err == nil {
			digits = append(digits, num)
		}
	}
	myNum := 0
	if len(digits) >= 2 {
		myNum = appendNums(digits[0], digits[len(digits)-1])
	} else {
		myNum = appendNums(digits[0], digits[0])
	}
	return myNum
}

func appendNums(a, b int) int {
	res, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		log.Fatal(err)
	}
	return res
}
