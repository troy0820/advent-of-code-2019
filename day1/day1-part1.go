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
	var answer int
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		answer += solve(input)
	}
	fmt.Println("answer is", answer)
}
