package main

import (
	"fmt"
)

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.24, 3))
}
