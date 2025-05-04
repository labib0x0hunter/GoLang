package main

import "fmt"

func reverse(s []int) []int {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1{
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func equal(a []int, b []int) bool {
	if len(a) != len(b){
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Function(){

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Function
	fmt.Println(len(s), cap(s))

	// // Reverse the array
	// reverse(s)
	// PrintSlice(s)

	// // Shift left by n element
	// // Also code for right shift [-]
	// n := 3
	// reverse(s[:n])
	// reverse(s[n:])
	// reverse(s)
	// PrintSlice(s)

	// // Check if two slices are equal, 
	// // s1 == s2 isn't acceptable
	// fmt.Println(equal([]int{1, 3, 4}, []int{1, 3, 4}))
	// fmt.Println(equal([]int{1, 2, 3}, []int{0, 1, 3}))

	// // copy function -> copy(dst, src)
	// // Remove range [2, 3] zero based index
	// s1 := []int{0, 1, 2, 3, 4, 5}
	// PrintSlice(s1[2:])
	// PrintSlice(s1[4:])
	// copy(s1[2 : ], s1[4 : ])
	// PrintSlice(s1[: len(s1) - (4 - 2)])

	// Insert at index i = 3
	s2 := []int{0, 1, 2, 4, 5, 6}
	s2 = append(s2, -1)
	PrintSlice(s2[4:])
	PrintSlice(s2[3:])
	copy(s2[4:], s2[3:])
	s2[3] = 3
	PrintSlice(s2)
}

