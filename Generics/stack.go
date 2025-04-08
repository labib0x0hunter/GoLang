package main

import (
	"fmt"
)

// Struct for stack
type Stack[T any] struct {
	stack []T
}

// Check if stack is empty
func (s *Stack[T]) empty() bool {
	return len(s.stack) == 0
}

// Returns the size of the stack
func (s *Stack[T]) size() int {
	return len(s.stack)
}

// Push an element on top
func (s *Stack[T]) push(element T){
	s.stack = append(s.stack, element)
}

// Remove top element
func (s *Stack[T]) pop() (bool) {
	length := len(s.stack)
	if length == 0 {
		return false
	}
	s.stack = s.stack[:length - 1]
	return true
}

// Returns top element without removing it
func (s *Stack[T]) peek() (T, bool) {
	length := len(s.stack)
	if length == 0 {
		var empty T
		return empty, false
	}
	last_element := s.stack[length - 1]
	return last_element, true
}

// Print Stack
func (s Stack[T]) String() string {
	return fmt.Sprintf("Stack: %v", s.stack)
}

// Clear the stack
func (s *Stack[T]) clear() {
	s.stack = nil
}

// Find strictly smaller value of arr[i] in left
// Returns the index
// -1 means no smaller value
func prev_smaller(arr []int) []int {
	n := len(arr)
	var st Stack[int]
	st.push(-1)
	index := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for {
			peek, ok := st.peek()
			if !ok || peek == - 1 || arr[i] > arr[peek] {
				break
			}
			st.pop()
		}
		peek, _ := st.peek()
		index = append(index, peek)
		st.push(i)
	}
	return index
}

// Find strictly greater value of arr[i] in left
// Returns the index
// -1 means no greater value
func prev_greater(arr []int) []int {
	n := len(arr)
	var st Stack[int]
	st.push(-1)
	index := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for {
			peek, ok := st.peek()
			if !ok || peek == - 1 || arr[i] < arr[peek] {
				break
			}
			st.pop()
		}
		peek, _ := st.peek()
		index = append(index, peek)
		st.push(i)
	}
	return index
}

func main() {
	
	var st Stack[int]

	st.push(10)

	fmt.Println(st.pop())
	fmt.Println(st.empty())


	fmt.Println(st.peek())
	st.push(20)
	st.push(30)
	fmt.Println(st.peek())

	fmt.Println(st)

	fmt.Println(st.size())
	st.clear()
	fmt.Println(st.empty())


	fmt.Println(prev_smaller([]int{1, 2, 5, 4, 3, 1}))
	fmt.Println(prev_greater([]int{1, 2, 5, 4, 3, 1}))
}
