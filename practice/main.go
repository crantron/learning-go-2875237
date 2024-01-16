package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := make([]int, 5)
	numbers[0] = 34
	numbers[1] = 32
	numbers[2] = 345
	numbers[3] = 346
	numbers[4] = 342
	fmt.Println(numbers)

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["CA"] = "California"
	states["AZ"] = "Arizona"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	slices.Sort(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)

	// Convert Slice of strings to a map
	var fruits []string = make([]string, 3)
	fruits[0] = "banana"
	fruits[1] = "orange"
	fruits[2] = "apple"

	var result map[string]float64 = converToMap(fruits)

	fmt.Println(result)
}

func converToMap(items []string) map[string]float64 {
	result := make(map[string]float64)
	var totalPrice float64 = 100
	var itemLen float64 = float64(len(items))
	var price float64 = totalPrice / itemLen

	for i := range items {
		result[items[i]] = price
	}

	return result
}

type Dog struct {
	Breed  string
	Weight int
}
