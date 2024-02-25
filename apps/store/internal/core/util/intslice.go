package util

func MaxIntSlice(values ...int) int {
	var match int
	for i, e := range values {
		if i == 0 || e > match {
			match = e
		}
	}

	return match
}

func MinIntSlice(values ...int) int {
	var match int
	for i, e := range values {
		if i == 0 || e < match {
			match = e
		}
	}

	return match
}
