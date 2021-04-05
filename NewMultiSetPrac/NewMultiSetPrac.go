package NewMultiSetPrac

func NewMultiSet() map[string]int {
	return map[string]int{}
}

func String(m map[string]int) string {
	result := "{ "

	for val, cnt := range m {
		for i := 0; i < cnt; i++ {
			result += val + " "
		}
	}

	result += "}"
	return result
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func Insert(m map[string]int, val string) {
	m[val] += 1
}

func Erase(m map[string]int, val string) {
	if m[val] != 0 {
		m[val] -= 1
	}
}
