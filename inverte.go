package main

import "fmt"

func main(){
	inverterArray(1, 2, 3 ,4)
}

func inverterArray(a, b, c, d int) {
	array := [4]int{a, b, c, d}
	inv := []int{}
	for i := len(array) - 1; i >= 0; i--{   
		inv = append(inv, array[i])
	}

	fmt.Println(inv)
}