package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum/float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		//parse each arg into a float because Args returns a slice of args
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("%#v\n", numbers)
	avg := average(numbers...)
	fmt.Println("average is ", avg)
}
