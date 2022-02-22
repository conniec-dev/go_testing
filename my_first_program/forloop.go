package main

import "fmt"

func main() {
	//The most basic type, with a single condition.
	i:= 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
for j := 7; j <= 9; j++ {
	fmt.Println(j)
}
for {
	fmt.Println("loop")
	break
}
}