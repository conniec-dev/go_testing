package main

import "fmt"

type IceCream struct {
	Flavor string
	Size   string
}

func main() {
	i := IceCream{
		Flavor: "chocolate",
		Size:   "Medium",
	}
	fmt.Println(i.Flavor)
}
