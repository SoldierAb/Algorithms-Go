package main

import "fmt"

func main(){

	arr := []int{10,1,8,2,3,7,9,4,5}

	for i:=1;i<len(arr);i++{

		if arr[i]<arr[i-1]{

			tmp:= arr[i]

			for j:=i;j>=0;j--{
				if j>0&&arr[j-1]>tmp{
					arr[j]=arr[j-1]
				}else{
					arr[j] = tmp
					break
				}
			}

		}

	}

	fmt.Println(arr)

}
