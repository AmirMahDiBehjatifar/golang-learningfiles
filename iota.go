package main

import "fmt"

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

func main() {
	fmt.Println("Tuesday is day number", Tuesday) // Output: Tuesday is day number 2
}
