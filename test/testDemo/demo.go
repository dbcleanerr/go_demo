package testDemo

func Add(a, b int) int {
	return a + b
}

func Square(a int) int {
	return a * a
}

func Sum(a, b int) int {
	sum := 0

	for i := a; i < b; i++ {
		sum += i
	}

	return sum
}