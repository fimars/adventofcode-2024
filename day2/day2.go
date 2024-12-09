package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day2_input")
	defer f.Close()

	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")

		if isSafe(strs) {
			count += 1
		}
	}

	fmt.Println(count)
}

func isSafe(strs []string) bool {
	f := ab
	current, next := 0, 0

	for i := 0; i < len(strs)-1; i++ {
		if i == 0 {
			current = conv(strs[i])
			next = conv(strs[i+1])
		} else {
			current = next
			next = conv(strs[i+1])
		}

		if next > current && i == 0 {
			f = bd
		}

		duration := f(current, next)

		if duration < 1 || duration > 3 {
			return false
		}
	}
	return true
}

func ab(a int, b int) int {
	return a - b
}
func bd(a int, b int) int {
	return b - a
}

func conv(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}
