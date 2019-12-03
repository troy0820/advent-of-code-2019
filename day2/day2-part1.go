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
	for i := 0; i < len(m); i = i + 4 {
		if m[i] == 1 {
			pos := m[i+1]
			pos2 := m[i+2]
			out := m[i+3]
			solution[out] = m[pos] + m[pos2]
			copy(m, solution)
		}
		if m[i] == 2 {
			pos := m[i+1]
			pos2 := m[i+2]
			out := m[i+3]
			solution[out] = m[pos] * m[pos2]
			copy(m, solution)
		}
		if m[i] == 99 {
			fmt.Println("stopping program: ")
			fmt.Println("solution:", solution)
		}
	}

	fmt.Println("length of solution array", len(solution))
}
