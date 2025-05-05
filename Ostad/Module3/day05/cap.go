package main

import "fmt"

// len(s) == cap(s), 
// In this situation append in s
// increases the capacity by 2 * len(s)
func Cap() {
	s := []int{1}

	fmt.Println(len(s), cap(s))
	s = append(s, 3)
	fmt.Println(len(s), cap(s))
	s = append(s, 4)
	fmt.Println(len(s), cap(s))
	s = append(s, 5)
	fmt.Println(len(s), cap(s))
	s = append(s, 6)
	fmt.Println(len(s), cap(s))
}
