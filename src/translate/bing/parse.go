package bing

import (
	"errors"
	"regexp"
)

var parsePattern, _ = regexp.Compile("translations\":.*?text\":\"(.*?)\",")

func parseResult(resultText string) (string, error) {
	matches := parsePattern.FindStringSubmatch(resultText)
	if 0 == len(matches) {
		return "", errors.New("none matches")
	}
	return matches[1], nil
}
