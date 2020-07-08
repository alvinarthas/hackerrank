package main

import "fmt"

/*
	Nyontek Version, i'm tired
*/
func main() {
	var numberRecs int

	fmt.Scan(&numberRecs)
	array := make([]int, numberRecs)
	//load the rectangles
	var val int
	for i := 0; i < numberRecs; i++ {
		fmt.Scan(&val)
		array[i] = val
	}

	maxArea := 0
	index := 0
	stack := []int{}

	for index < numberRecs {
		if len(stack) == 0 || array[stack[0]] <= array[index] {
			//push onto stack
			newStack := []int{index}
			newStack = append(newStack, stack...)
			stack = newStack
			//move to next index
			index++
		} else {
			//get the earlier edge of the rectangle
			lastIndex := stack[0]
			stack = stack[1:]
			//the length is either:
			//1. the last index
			//2. (index - (current top of stack-1)
			length := index
			if len(stack) > 0 {
				length = (index - stack[0] - 1)
			}
			//area == length * size at lastIndex
			area := array[lastIndex] * length
			if maxArea < area {
				maxArea = area
			}
		}
	}

	for len(stack) > 0 {
		top := stack[0]
		stack = stack[1:]
		length := index
		if len(stack) > 0 {
			length = (index - stack[0] - 1)
		}
		area := array[top] * length
		if maxArea < area {
			maxArea = area
		}
	}
	fmt.Println(maxArea)
}
