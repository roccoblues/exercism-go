package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

func (l IntList) Length() int {
	var n int
	for range l {
		n++
	}
	return n
}

func (l IntList) Reverse() IntList {
	out := make(IntList, l.Length())
	for i, v := range l {
		out[l.Length()-1-i] = v
	}
	return out
}

func (l IntList) Append(input IntList) IntList {
	out := make(IntList, l.Length()+input.Length())
	for i, v := range l {
		out[i] = v
	}
	for i, v := range input {
		out[i+l.Length()] = v
	}
	return out
}

func (l IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		l = l.Append(v)
	}
	return l
}

func (l IntList) Map(f unaryFunc) IntList {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}

func (l IntList) Foldr(f binFunc, x int) int {
	for i := range l {
		x = f(l[l.Length()-1-i], x)
	}
	return x
}

func (l IntList) Foldl(f binFunc, x int) int {
	for _, v := range l {
		x = f(x, v)
	}
	return x
}

func (l IntList) Filter(f predFunc) IntList {
	out := IntList{}
	for _, v := range l {
		if f(v) {
			out = out.Append(IntList{v})
		}
	}
	return out
}
