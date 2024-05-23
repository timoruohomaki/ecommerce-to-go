package register

import "strings"

func IsEmailValid(email string) bool {
	// @
	if strings.Count(email, "@") != 1 {
		return false
	}

	// empty
	segments := strings.Split(email, "@")
	addr, domain := segments[0], segments[1]

	if len(addr) == 0 || len(domain) == 0 {
		return false
	}

	// dot
	if !strings.Contains(domain, ".") {
		return false
	}

	// all good

	return true

}
