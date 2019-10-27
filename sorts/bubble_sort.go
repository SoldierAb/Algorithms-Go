package main

import "fmt"

func Swap(arr []int,i,j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j]= tmp
}


func BubbleSort (arr []int){

	swaped:=true
	for swaped{
		swaped = false
		for i:= 0;i<len(arr)-1;i++{
			if(arr[i]>arr[i+1]){
				Swap(arr,i,i+1)
				swaped=true
			}
		}
	}

}


func main(){
	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	BubbleSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}

