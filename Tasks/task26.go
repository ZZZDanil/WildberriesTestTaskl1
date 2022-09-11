package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "zZ"
	fmt.Println(s1, " ", IsUniq(&s1))
	fmt.Println(s2, " ", IsUniq(&s2))
	fmt.Println(s3, " ", IsUniq(&s3))
	fmt.Println(s4, " ", IsUniq(&s4))
}

// прохождение по всей строке
// карта заполняется рунами
// проверка руны на наличие в карте
// уникальность - если в карте только одна руна по содержимому
func IsUniq(str *string) bool {
	m := make(map[rune]int)
	for _, v := range *str {
		v = unicode.ToLower(v)
		_, ok := m[v]
		if ok {
			return false
		} else {
			m[v] = 0
		}
	}
	return true
}
