package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := "input.txt"
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	sum := 0
	c, _, _ := reader.ReadRune()
	first, _ := strconv.Atoi(string(c))
	cur := first
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				break
			}
		}

		next, _ := strconv.Atoi(string(c))
		if cur == next {
			sum += cur
		}

		cur = next
	}

	if cur == first {
		sum += cur
	}

	fmt.Printf("sum = %d", sum)
}
