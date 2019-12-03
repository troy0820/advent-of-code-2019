package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertToInts(s []string) []int {
	n := []int{}
	for _, v := range s {
		m, _ := strconv.Atoi(v)
		n = append(n, m)
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), ",")
	}
	m := ConvertToInts(input)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			solution := make([]int, len(m))
			copy(solution, m)
			solution[1], solution[2] = noun, verb
			for i := 0; i < len(m); i = i + 4 {
				if solution[i] == 1 {
					solution[solution[i+3]] = solution[solution[i+1]] + solution[solution[i+2]]
				}
				if solution[i] == 2 {
					solution[solution[i+3]] = solution[solution[i+1]] * solution[solution[i+2]]
				}
				if solution[0] == 19690720 {
					fmt.Printf("Noun: %d Verb: %d, Solution: %d \n", noun, verb, 100*noun+verb)
					break
				}
			}
		}
	}
}
