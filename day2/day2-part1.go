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
	solution := make([]int, len(m))
	copy(solution, m)
	solution[1], solution[2] = 12, 2
	for i := 0; i < len(m); i = i + 4 {
		if m[i] == 1 {
			solution[m[i+3]] = m[m[i+1]] + m[m[i+2]]
			if m[i+4] == 99 {
			} else {
				copy(m, solution)
			}
		}
		if m[i] == 2 {
			solution[m[i+3]] = m[m[i+1]] * m[m[i+2]]
			if m[i+4] == 99 {
			} else {
				copy(m, solution)
			}
		}
		if m[i] == 99 {
			fmt.Println("solution:", solution[0])
		}
	}
}
