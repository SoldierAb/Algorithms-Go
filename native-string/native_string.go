package main

import "fmt"

func MatchString(str,mat string) [] int{
	var pos []int
	for i:=0;i<len(str);i++{
		match := true
		for j:=0;j<len(mat);j++{
			if str[i+j]!=mat[j]{
				match = false
				break
			}
		}
		if match{
			pos = append(pos,i)
		}
	}
	return pos
}

func main(){
	text := `asdasfdsafsadfdfsdjkjionnkhioysdflkeojmasfdivcnnriehg`

	match_str := `as`

	fmt.Println(MatchString(text,match_str))
}


