package extensions

func Contains(s []interface{}, e int) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}
