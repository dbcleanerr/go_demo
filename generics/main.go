package main

// golang 1.8 开始提供generics(泛型)

import "fmt"

type Number interface {
	int64 | float64
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

func SumIntsOrFloats[V int64 | float64](m map[string]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func SumNumber[V Number](m map[string]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	m1 := map[string]int64{
		"a": 10,
		"b": 20,
	}

	m2 := map[string]float64{
		"a": 15.1,
		"b": 25.2,
	}

	fmt.Println(SumInts(m1))
	fmt.Println(SumFloats(m2))

	s1 := SumIntsOrFloats(m1)
	s2 := SumIntsOrFloats(m2)

	fmt.Printf("s1 %T\n", s1)
	fmt.Printf("s2 %T\n", s2)

	fmt.Println(SumNumber(m1))
	fmt.Println(SumNumber(m2))
}
