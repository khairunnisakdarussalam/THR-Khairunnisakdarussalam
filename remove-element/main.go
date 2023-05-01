package main

import "fmt"

// a := [4]int{1, 2, 3, 4} // "a" has type [4]int (array of 4 ints)
// x := a[:]   // "x" has type []int (slice of ints) and length 4
// y := a[:2]  // "y" has type []int, length 2, values {1, 2}
// z := a[2:]  // "z" has type []int, length 2, values {3, 4}
// m := a[1:3] // "m" has type []int, length 2, values {2, 3}

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	var last = []interface{}{}

	if position == "left" {
		last = append(last, arr[1:]...)
	} else if position == "right" {
		last = append(last, arr[:len(arr)-1]...)
	} else {
		//
	}
	return last
}

func main() {
	var arr = []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(arr, "left"))
}
