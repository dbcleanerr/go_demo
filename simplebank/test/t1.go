package main

import "fmt"

type ss struct {
	address string
	port    int
}

func main() {
	s1 := ss{
		address: "0.0.0.0",
		port:    0,
	}

	fmt.Printf("%#v\n", s1)

	s2 := &ss{
		address: "1.1.1.1",
		port:    0,
	}
	fmt.Printf("%#v\n", s2)

}
