package typoglycemia

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func TypoGlycemia(s string) {
	sl := strings.Split(s, " ")
	fmt.Println(sl)
	data := shuffle(sl)
	fmt.Println(data)
}

func shuffle(sl []string) []string {
	rand.Seed(time.Now().UnixNano())

	first := sl[0]
	sl = sl[1:]
	last := sl[len(sl)-1]
	sl = sl[:len(sl)-1]

	n := len(sl)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		sl[i], sl[j] = sl[j], sl[i]
	}

	var res []string
	res = append(res, first)
	res = append(res, sl...)
	res = append(res, last)

	return res
}
