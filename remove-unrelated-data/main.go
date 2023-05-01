package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	_, isFound := dataMap[key]
	if isFound {
		delete(dataMap, key)
	} else {
		fmt.Println("Not Found!")
	}
	return dataMap
}

func main() {
	// data := make(map[string]interface{})
	data := map[string]interface{}{"run": "lari", "jump": "loncat", "swim": "berenang"}
	fmt.Println(removeUnrelated(data, "jump"))
}
