package main

import (
	"fmt"
	"math/rand"
)

func QuickSort (arr []int) []int{
	if len(arr)<=1{
		return arr
	}

	midian := arr[rand.Intn(len(arr))]

	leftArr := make([]int,0,len(arr))
	rightArr := make([]int,0,len(arr))
	middleArr := make([]int,0,len(arr))

	for _,val := range arr{
		switch {
		case val<midian:
			leftArr = append(leftArr,val)
		case val == midian:
			middleArr = append(middleArr,val)
		case val >midian:
			rightArr = append(rightArr,val)
		}
	}

	leftArr = QuickSort(leftArr)
	rightArr = QuickSort(rightArr)

	leftArr = append(leftArr,middleArr...)
	leftArr = append(leftArr,rightArr...)

	return leftArr

}


func RandomArr (l int)[]int{
	arr := make([]int,l)
	for i:=0 ;i<=l-1;i++{
		arr[i] = rand.Intn(l)
	}
	return arr
}


func main(){
	arr := RandomArr(10)

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", QuickSort(arr))
}