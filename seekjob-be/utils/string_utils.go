package utils

import "strings"

func ConvertStringSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "%20")
}

func ConstructAPIPath(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	ret := ""
	for i := 0; i < len(strings); i++ {
		ret += strings[i]
		if i == len(strings)-1 {
			break
		}
		ret += "/"
	}
	return ret
}

func ConstructAPIQuery(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	ret := ""
	operators := []string{"=", "&"}
	for i := 0; i < len(strings); i++ {
		ret += strings[i]
		if i == len(strings)-1 {
			break
		}
		ret += operators[i%2]
	}
	return ret
}
