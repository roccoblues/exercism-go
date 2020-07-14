package flatten

func Flatten(input interface{}) []interface{} {
	out := []interface{}{}
	for _, v := range input.([]interface{}) {
		if v != nil {
			subList, ok := v.([]interface{})
			if ok {
				out = append(out, Flatten(subList)...)
			} else {
				out = append(out, v)
			}
		}
	}
	return out
}
