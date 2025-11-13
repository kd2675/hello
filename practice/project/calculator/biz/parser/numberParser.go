package parser

import (
	"regexp"
	"strconv"
	"strings"
)

type NumberParser interface {
	ParseNumbers(body string, delimiters []string) ([]int64, error)
}

type numberParserImpl struct{}

func NewNumberParser() NumberParser {
	return &numberParserImpl{}
}

func (p *numberParserImpl) ParseNumbers(body string, delimiters []string) ([]int64, error) {
	if body == "" {
		return []int64{}, nil
	}

	var escapedDelimiters []string
	for _, d := range delimiters {
		escaped := regexp.QuoteMeta(d)
		escapedDelimiters = append(escapedDelimiters, escaped)
	}

	pattern := strings.Join(escapedDelimiters, "|")
	re := regexp.MustCompile(pattern)

	tokens := re.Split(body, -1)

	var numbers []int64
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		num, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}
