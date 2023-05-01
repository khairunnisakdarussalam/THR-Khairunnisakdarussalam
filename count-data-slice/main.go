package main

import "fmt"

func howManyElements(data []any) int {
	//your code here!
	return len(data)
}

func main() {
	var data = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(data))
}
