package main

import "fmt"

// Define a generic Stack
type Stack[T any] struct {
	items []T
}

// Push adds an element to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack[T]) Pop() (T) {
	var zero T

	if len(s.items) == 0 {
		return zero
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	var zero T

	if len(s.items) == 0 {
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty checks if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns number of elements
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func hasValidPath(grid [][]int) bool {
	prev := [2]int{}
	coordinates := [2]int{}
	temp := [2]int{}
	expected := [2][2]int{}
	incoming := [2][2]int{}
	maxX := len(grid[0])
	maxY := len(grid)
	validCheck := false
	connectCheck := false
	count := 0
	stack := Stack[[2]int]{}

	for {
		fmt.Print("coordinates: ")
		fmt.Println(coordinates)
		if coordinates == [2]int{maxX - 1, maxY - 1} {
			return true
		}

		expected = checkConnect(grid[coordinates[1]][coordinates[0]], coordinates)
		fmt.Print("expected: ")
		fmt.Println(expected)
		for _, value := range expected {
			if checkValid(value, maxX, maxY) && value != prev {
				if count == 0 {
					stack.Push(value)
					validCheck = true
				} else {
					temp = value
					validCheck = true
				}
			}
		}
		fmt.Print("stack: ")
		fmt.Println(stack)

		if count == 0 {
			temp = stack.Pop()
		}

		fmt.Print("temp: ")
		fmt.Println(temp)
		if validCheck{
			validCheck = false
		} else if !stack.IsEmpty() {
			temp = stack.Pop()
			coordinates = [2]int{0, 0}
			validCheck = false
		} else {
			return false
		}
		fmt.Print("temp: ")
		fmt.Println(temp)

		incoming = checkConnect(grid[temp[1]][temp[0]], temp)
		fmt.Print("incoming: ")
		fmt.Println(incoming)
		for _, value := range incoming {
			if value == coordinates {
				prev = coordinates
				coordinates = temp
				connectCheck = true
			}
		}

		fmt.Print("coordinates: ")
		fmt.Println(coordinates)
		if connectCheck {
			connectCheck = false
		} else if !stack.IsEmpty() {
			coordinates = stack.Pop()
			prev = [2]int{0, 0}
			connectCheck = false
			fmt.Print("stack: ")
			fmt.Println(stack)
		} else {
			return false
		}
		count++
	}
}

func checkValid(nums [2]int, maxX int, maxY int) bool {
	if nums[0] >= 0 && nums[0] < maxX && nums[1] >= 0 && nums[1] < maxY {
		return true
	} else {
		return false
	}
}

func checkConnect(num int, coordinates [2]int) [2][2]int {
	connects := [2][2]int{}

	switch num {
	case 1:
		connects[0] = directionConvert(coordinates, "left")
		connects[1] = directionConvert(coordinates, "right")
	case 2:
		connects[0] = directionConvert(coordinates, "up")
		connects[1] = directionConvert(coordinates, "down")
	case 3:
		connects[0] = directionConvert(coordinates, "left")
		connects[1] = directionConvert(coordinates, "down")
	case 4:
		connects[0] = directionConvert(coordinates, "down")
		connects[1] = directionConvert(coordinates, "right")
	case 5:
		connects[0] = directionConvert(coordinates, "up")
		connects[1] = directionConvert(coordinates, "left")
	case 6:
		connects[0] = directionConvert(coordinates, "up")
		connects[1] = directionConvert(coordinates, "right")
	}

	return connects
}

func directionConvert(coordinates [2]int, direction string) [2]int {
	result := [2]int{}

	switch direction {
	case "up":
		result[0] = coordinates[0]
		result[1] = coordinates[1] - 1
	case "left":
		result[0] = coordinates[0] - 1
		result[1] = coordinates[1]
	case "right":
		result[0] = coordinates[0] + 1
		result[1] = coordinates[1]
	case "down":
		result[0] = coordinates[0]
		result[1] = coordinates[1] + 1
	}

	return result
}