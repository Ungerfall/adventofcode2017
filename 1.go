package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "input.txt"
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	sum := 0
	c, _, _ := reader.ReadRune()
	var num int
	first, _ := strconv.Atoi(string(c))
	fmt.Printf("first = %d\n", first)
	num = first
	for {
		c, _, err := reader.ReadRune()
		fmt.Printf(string(c))
		if err != nil && err == io.EOF {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
				break
			}
		}

		next, _ := strconv.Atoi(string(c))
		if num == next {
			sum += num
		}

		num = next
	}

	fmt.Printf("\nlast = %d", num)
	if num == first {
		sum += num
	}

	fmt.Printf("sum = %d", sum)
}
