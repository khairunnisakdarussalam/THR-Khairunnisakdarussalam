package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	//your code here!
	splitdata := strings.Split(data, ",")

	newname := string(splitdata[0])
	newage, _ := strconv.Atoi(splitdata[1])
	newadress := string(splitdata[2])

	return User{Name: newname, Age: newage, Address: newadress}
}

func main() {
	fmt.Println(ConvertData("Edo,20,Jakarta"))
}
