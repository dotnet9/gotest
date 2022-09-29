package main

import (
	"fmt"
	"log"

	"dotnet9.com/greetings"
)

type Number interface {
	int64 | float64
}

func main() {

	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}
	floats := map[string]float64{
		"first": 359.98,
		"second": 26.99,
	}
	fmt.Printf("Non-Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	//fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Dotnet9", "C#", "Go"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	} 
	
	fmt.Println(messages)
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}