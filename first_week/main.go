package main

import "fmt"

func main() {
	slice := []int{111, 2, 3, 4, 5}
	tt := sliDelete(1, slice)
	fmt.Println(tt)
}
