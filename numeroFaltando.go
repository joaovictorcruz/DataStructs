package main

import (
	"fmt"
	"sort"
) 

func numeroFaltando(nums []int) int{

	arrLength := len(nums)
	sum := arrLength * (arrLength + 1 ) / 2

	for _, value :=  range nums{
		sum -= value
	}

	fmt.Println(sum)
	return sum
}

func numeroFaltandoNaUnha(nums []int) int {
    n := len(nums)
    sort.Ints(nums)

    esperado := make([]int, n+1)
    for i := 0; i <= n; i++ {
        esperado[i] = i
    }

    for i := 0; i < n; i++ {
        if nums[i] != esperado[i] {
            return esperado[i]
        }
    }

	fmt.Println(esperado[n])
    return esperado[n]
}
