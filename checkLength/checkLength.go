package main

import "fmt"

func checkLength(ss *[]string) {
	if len(*ss) != 1 {
		*ss = nil
	}
}

func main() {
	data1 := []string{"ABC", "abc"}
	data2 := []string{"ABC"}

	checkLength(&data1)
	checkLength(&data2)
	fmt.Println(data1)
	fmt.Println(data2)
}


