package regexhelpers

import "regexp"

// FindAllStringMatches performs a regexp match for regexCondition on baseString
func FindAllStringMatches(baseString string, regexCondition string) []string {
	condition, _ := regexp.Compile(regexCondition)
	return condition.FindAllString(string(baseString), -1)
}
