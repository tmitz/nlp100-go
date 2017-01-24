package set

func Ngram(s string, n int) []string {
	var res []string

	for i := 0; i < len(s)-n+1; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func Intersection(x, y []string) []string {
	var res []string

	m := make(map[string]struct{})
	for _, xword := range x {
		for _, yword := range y {
			if xword == yword {
				m[xword] = struct{}{}
			}
		}
	}
	for k := range m {
		res = append(res, k)
	}

	return res
}

func Difference(x, y []string) []string {
	var res []string
	m := make(map[string]struct{})
	for _, xword := range x {
		for _, yword := range y {
			if xword == yword {
				continue
			} else {
				m[xword] = struct{}{}
			}
		}
	}
	for k := range m {
		res = append(res, k)
	}

	return res
}

func Union(x, y []string) []string {
	var res []string
	m := make(map[string]struct{}, len(x)+len(y))
	for _, xword := range x {
		m[xword] = struct{}{}
	}
	for _, yword := range y {
		m[yword] = struct{}{}
	}
	for k := range m {
		res = append(res, k)
	}

	return res
}

func IncludeBiGram(l []string, substr string) bool {
	for _, word := range l {
		if word == substr {
			return true
		}
	}
	return false
}
