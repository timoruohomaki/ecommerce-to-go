package register

func IsPasswordValid(pass string, length int) bool {
	// length
	if len(pass) < length {
		return false
	}
	// rules
	var lower, upper, numeric bool

	for _, char := range pass {
		switch {
		case char >= 'a' && char <= 'z':
			lower = true
		case char >= 'A' && char <= 'Z':
			upper = true
		case char >= '0' && char <= '9':
			numeric = true
		}
	}

	return lower && upper && numeric

}
