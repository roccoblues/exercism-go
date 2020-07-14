package secret

var events []string = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	reverse := code&(1<<4) != 0
	length := len(events)

	e := []string{}
	for i := 0; i < length; i++ {
		n := i
		if reverse {
			n = length - 1 - i
		}
		if code&(1<<n) != 0 {
			e = append(e, events[n])
		}
	}

	return e
}
