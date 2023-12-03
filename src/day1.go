package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	toSum := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		digits := []int{}
		for i := 0; i < len(line); i++ {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				digits = append(digits, num)
			}
		}
		myNum := 0
		if len(digits) >= 2 {
			myNum = appendNums(digits[0], digits[len(digits)-1])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			myNum = appendNums(digits[0], digits[0])
		}
		toSum = append(toSum, myNum)
	}

	total := 0
	for _, s := range toSum {
		total = total + s
	}
	fmt.Println(total)

}

func appendNums(a, b int) int {
	res, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		log.Fatal(err)
	}
	return res
}
