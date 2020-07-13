package sublist

type Relation string

func Sublist(a, b []int) Relation {
	var relation Relation

	if len(a) == len(b) {
		relation = "equal"
	} else if len(a) < len(b) {
		relation = "sublist"
	} else {
		relation = "superlist"
		a, b = b, a
	}

	ai := 0
	bi := 0
	for {
		if ai > len(a)-1 {
			return relation
		}
		if bi > len(b)-1 {
			return "unequal"
		}

		if a[ai] == b[bi] {
			ai++
			bi++
			continue
		}

		bi -= ai
		ai = 0
		bi++
	}
}
