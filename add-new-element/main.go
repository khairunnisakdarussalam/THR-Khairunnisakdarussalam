package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!
	var arr = []int{}

	switch position {
	case "up":
		arr = append(arr, newData)
		arr = append(arr, data...)
		return arr
	case "down":
		arr = append(arr, data...)
		arr = append(arr, newData)
		return arr
	default:
		return data
	}
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(AddElement(arr, 6, "down"))
}
