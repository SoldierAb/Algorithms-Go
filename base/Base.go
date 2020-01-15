package base

import (
	"math/rand"
	"time"
)

type  Base struct{}


func (b *Base) Swap(arr []interface{},i,j int){
	tmp:= arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}


func (b *Base) RandomInt(count int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(count)
}