package parser

import (
	"errors"
	"regexp"
	"strings"
)

type DelimiterParserImpl struct{}

func NewDelimiterParserImpl() DelimiterParser {
	return &DelimiterParserImpl{}
}

func (d *DelimiterParserImpl) ParseDelimiters(input string) ([]string, error) {
	if err := d.validateInput(input, "parseDelimiters"); err != nil {
		return nil, err
	}

	if !d.hasCustomDelimiter(input) {
		return defaultDelimiters, nil
	}

	customDelimiter, err := d.extractCustomDelimiter(input)
	if err != nil {
		return nil, err
	}

	escapedCustom := regexp.QuoteMeta(customDelimiter)
	return []string{escapedCustom, ",", ":"}, nil
}

func (d *DelimiterParserImpl) ExtractBody(input string) (string, error) {
	if err := d.validateInput(input, "extractBody"); err != nil {
		return "", err
	}

	if !d.hasCustomDelimiter(input) {
		return input, nil
	}

	delimiterEndIndex, err := d.getDelimiterEndIndex(input)
	if err != nil {
		return "", err
	}

	return input[delimiterEndIndex+suffixLength:], nil
}

func (d *DelimiterParserImpl) validateInput(input string, methodName string) error {
	if input == "" {
		return errors.New(methodName + " input is empty")
	}
	return nil
}

func (d *DelimiterParserImpl) hasCustomDelimiter(input string) bool {
	return strings.HasPrefix(input, customDelimiterPrefix)
}

func (d *DelimiterParserImpl) extractCustomDelimiter(input string) (string, error) {
	delimiterEndIndex, err := d.getDelimiterEndIndex(input)
	if err != nil {
		return "", err
	}
	return input[prefixLength:delimiterEndIndex], nil
}

func (d *DelimiterParserImpl) getDelimiterEndIndex(input string) (int, error) {
	index := strings.Index(input, customDelimiterSuffix)
	if index < 0 {
		return -1, errors.New("커스텀 구분자 형식이 올바르지 않습니다. (//구분자\\n 형식이어야 합니다)")
	}
	return index, nil
}
