package main

import "fmt"

func main() {
	// Input := []int{1, -2, 0, 3}
	Input := []int{3, -1, -1, 4, 3, -1} 

	a := 0
	for i := range Input{
		v := Input[i:]
		s := Sum(v)
		if s > a {
			a = s
		} 
	}
	
	r := Reverse(Input)
	for _, v := range r{
		if  v < 0 {
			a -= v
		} else {
			break
		}
		
	}


	fmt.Print(a)
}

func ArrayChallenge(arr []int) int {
	return 0
}

// func checkInit(a []int) int{
// 	s := Sum(a)
// }

func Sum(a []int) int{
	if len(a) == 1 {
		// fmt.Println(a)
		return a[0]
	}

	return Sum(a[:1]) + Sum(a[1:])
}

func Reverse(a []int) (r []int) {
	if len(a) == 1 {
		return append(r, a[0])
	}
	return append(a[len(a)-1:],a[:len(a)-1]...)
}