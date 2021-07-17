package main

import (
	"fmt"
)

func main() {
	f := func(a []string) ([]string, string) {
		return a[1:], a[0] // 2番目以降の配列と1番目の要素を返す
	}

	m := []string{
		"one",
		"two",
		"three",
	}

	s := ""
	fmt.Println(m)
	for len(m) > 0 {
		m, s = f(m)
		fmt.Println(s+"->", m)
	}
}
