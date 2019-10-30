package main

import "fmt"

func main(){
	arr := []int{1,5,22,11,10,3,2,7,18}

	var min int  = 0
	var tmp int

	for i:= 0;i<len(arr);i++{
		min = i
		for j:= i+1;j<len(arr);j++{
			if arr[j]<arr[min]{
				min = j
			}
		}
		tmp  =arr[i]
		arr[i] = arr[min]
		arr[min]=  tmp
	}

	fmt.Println("Selected Sorted arr:  ",arr)
}
