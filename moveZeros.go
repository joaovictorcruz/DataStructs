package main

import "fmt"

func moveZeros(arr []int) []int{

	a := 0
	
	for i:= 0; i < len(arr); i++{
		if arr[i] != 0 && arr[a] == 0{ 
			arr[i], arr[a] = arr[a], arr[i]
		}

		if arr[a] != 0{
			a++
		}
	}

	fmt.Println(arr)
	return arr
}