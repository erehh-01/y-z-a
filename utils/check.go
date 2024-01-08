package utils

import "regexp"

func IsValidURL(input string) bool {
	regexPattern := `^(http|https):\/\/[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)+([\/?].*)?$`

	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}
