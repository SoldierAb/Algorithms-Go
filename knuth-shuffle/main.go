package main

import (
	"Algorithms-Go/base"
	"fmt"
)

type Knuth struct {
	base.Base
}

func (k *Knuth) KnuthShuffle(arr []interface{}, count int) []interface{} {
	for i := len(arr) - 1; i > 0; i-- {
		ranInt := k.RandomInt(i)
		k.Swap(arr, i, ranInt)
	}
	fmt.Println(arr)
	return arr[len(arr)-count:]
}

func main() {
	k := &Knuth{}
	arr := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(k.KnuthShuffle(arr, 2))
}
