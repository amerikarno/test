package main

import (
	"fmt"
	"strings"
)


func main() {
	// a := "abcabcabcabc"
	// a := "abcaacabcabcabc"
	a := "abcdabcdabcdabcd"

	d := Data(a)
	fmt.Println(d)
}

func Data (s1 string) (s string) {

	ss := strings.Split(s1,"")
	lss := LenOfSlice(ss)

	ans := "-1"
	for _, l := range lss {
		r := Slice(ss, l)
		d := Duplicate(r)
		if d != "-1" {
			ans = d
		}
	}

	return ans

}

func Duplicate(strSlice []string) string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list[0]
}

func Slice(s []string, l int) (ss []string) {
	if len(s) == 2*l {

		if strings.Join(s[0:l],"") == strings.Join(s[l:],"") {
			return append(ss, strings.Join(s[0:l],""))
			} else {
				return []string{"-1"}
			}
		}

	return append(Slice(s[0:2*l],l), Slice(s[l:],l)...)
}

func LenOfSlice(s []string) (si []int){
	ml := len(s)/2

	for i := 2; i <= ml; i++ {
		if len(s)%i == 0 {
			si = append(si, i)
		}
	}
	return si
}