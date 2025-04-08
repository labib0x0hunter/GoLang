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
// if stack is empty then return T, false
// else return top element , true
func (s *Stack[T]) pop() (T, bool) {
	length := len(s.stack)
	if length == 0 {
		var empty T
		return empty, false
	}
	last_element := s.stack[length - 1]
	s.stack = s.stack[:length - 1]
	return last_element, true
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
}
