package main

import "fmt"

func main(){
	inverterArray(1, 2, 3 ,4)

	arr := []int{1,2, 3, 4}
	inverteArray2(arr)
}

func inverterArray(a, b, c, d int) {
	array := [4]int{a, b, c, d}
	inv := []int{}
	for i := len(array) - 1; i >= 0; i--{     
		inv = append(inv, array[i])
	}

	fmt.Println(inv)
}

  //sem append
func inverteArray2(arr [] int) []int{
	start := 0
	end := len(arr) - 1
	

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	fmt.Println(arr)
	return  arr
}