package repeat

func Repeat(ch string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += ch
	}
	return result
}
