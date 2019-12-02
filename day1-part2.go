package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(a int) int {
	return a/3 - 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	answers := []int{}
	var answer int
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		answer = solve(input)
		var added int
		for answer/3 > 2 {
			answer = solve(answer)
			added += answer
		}
		firstanswer := solve(input)
		answers = append(answers, added+firstanswer)
	}
	var solution int
	for _, v := range answers {
		solution += v
	}
	fmt.Println("Solution", solution)

}
