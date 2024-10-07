package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1(a);2(b);3(c)"
	var res []string
	split := strings.Split(s, ";")
	fmt.Println(split)
	for i := 0; i < len(split); i++ {
		s2 := strings.Split(split[i], "(")[1]
		s2 = strings.ReplaceAll(s2, ")", "")
		res = append(res, s2)
	}
	fmt.Println(strings.Join(res, ";"))

}
