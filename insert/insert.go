package main

import "fmt"

func main() {
	org := []string{"Test1", "Test2", "Test3"}
	add := "Test4"

	fmt.Println(Insert(org, add))
}

func Insert(origin []string, ins string) (s []string) {
	return append([]string{ins}, origin...)
}