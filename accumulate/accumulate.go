package accumulate

func Accumulate(in []string, f func(string) string) []string {
	for i, v := range in {
		in[i] = f(v)
	}
	return in
}
