package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Ngram(input interface{}, n int) []string {
	var res []string

	switch reflect.ValueOf(input).Kind() {
	case reflect.String:
		str := input.(string)
		for i := 0; i < len(str)-n+1; i++ {
			res = append(res, str[i:i+n])
		}
	case reflect.Slice:
		arr := input.([]string)
		for i := 0; i < len(arr)-n+1; i++ {
			res = append(res, strings.Join(arr[i:i+n], "-"))
		}
	default:
		panic("not supported type...")
	}
	return res
}

func main() {
	s := "I am an NLPer"
	fmt.Println(Ngram(s, 2))
	arr := strings.Split(s, " ")
	fmt.Println(Ngram(arr, 2))
}
