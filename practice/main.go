package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(value1 string, value2 string) float64 {

	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 32)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 32)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	return float1 + float2
}

func main() {
	value1 := "0"
	value2 := "0"
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter First Float: ")
	input, err := reader.ReadString('\n')
	value1 = input
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter Second Float: ")
	input, err = reader.ReadString('\n')
	value2 = input
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sum: ", calculate(value1, value2))
}
