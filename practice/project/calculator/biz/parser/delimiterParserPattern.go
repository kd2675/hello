package parser

import (
	"errors"
	"regexp"
	"strings"
)

var customPattern = regexp.MustCompile(`^//(.+)\\n(.*)$`)

type DelimiterParserPattern struct{}

func NewDelimiterParserPattern() DelimiterParser {
	return &DelimiterParserPattern{}
}

func (d *DelimiterParserPattern) ParseDelimiters(input string) ([]string, error) {
	if input == "" {
		return nil, errors.New("parseDelimiters input is empty")
	}

	matches := customPattern.FindStringSubmatch(input)
	if len(matches) > 0 {
		custom := matches[1]
		if custom == "" {
			return nil, errors.New("커스텀 구분자가 비어있습니다")
		}
		escapedCustom := regexp.QuoteMeta(custom)
		return []string{escapedCustom, ",", ":"}, nil
	}

	if strings.HasPrefix(input, "//") {
		return nil, errors.New("커스텀 구분자 형식이 올바르지 않습니다. (//구분자\\n본문)")
	}

	return defaultDelimiters, nil
}

func (d *DelimiterParserPattern) ExtractBody(input string) (string, error) {
	if input == "" {
		return "", errors.New("extractBody input is empty")
	}

	matches := customPattern.FindStringSubmatch(input)
	if len(matches) > 0 {
		body := matches[2]
		if body == "" {
			return "", errors.New("본문이 비어있습니다")
		}
		return body, nil
	}

	if strings.HasPrefix(input, "//") {
		return "", errors.New("커스텀 구분자 형식이 올바르지 않습니다. (//구분자\\n본문)")
	}

	return input, nil
}
