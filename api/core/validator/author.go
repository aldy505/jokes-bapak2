package validator

import (
	"regexp"
	"strings"
)

func ValidateAuthor(author string) bool {
	if len(author) > 200 {
		return false
	}

	split := strings.Split(author, " ")
	if strings.HasPrefix(split[0], "<") && strings.HasSuffix(split[0], ">") {
		return false
	}
	if !strings.HasPrefix(split[len(split)-1], "<") && !strings.HasSuffix(split[len(split)-1], ">") {
		return false
	}

	email := strings.Replace(split[len(split)-1], "<", "", 1)
	email = strings.Replace(email, ">", "", 1)
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return pattern.MatchString(email)
}
