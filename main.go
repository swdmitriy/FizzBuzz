package main

import (
	"fmt"
	"strconv"
)

const (
	FIZZ string = "Fizz"
	BUZZ string = "Buzz"
)

func main() {
	for index := 1; index <= 100; index++ {
		fmt.Println(FizzBuzz(index))
	}
}

func FizzBuzz(input int) string {
	if input%15 == 0 {
		return FIZZ + BUZZ
	}
	if input%3 == 0 {
		return FIZZ
	}
	if input%5 == 0 {
		return BUZZ
	}
	return strconv.Itoa(input)
}