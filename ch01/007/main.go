package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(StrTemplate(12, "気温", 22.4))
}

func StrTemplate(x, y, z interface{}) string {
	return fmt.Sprintf("%s時の%sは%s", anytoa(x), anytoa(y), anytoa(z))
}

func anytoa(arg interface{}) string {
	var res string

	switch x := arg.(type) {
	case int:
		res = strconv.Itoa(x)
	case float32, float64:
		res = fmt.Sprint(x)
	case string:
		res = x
	case nil:
		res = "nil"
	case bool:
		if x {
			res = "true"
		} else {
			res = "false"
		}
	}

	return res
}
