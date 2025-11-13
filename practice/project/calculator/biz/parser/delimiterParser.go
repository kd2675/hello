package parser

type DelimiterParser interface {
	ParseDelimiters(input string) ([]string, error)
	ExtractBody(input string) (string, error)
}

const (
	customDelimiterPrefix = "//"
	customDelimiterSuffix = "\\n"
	prefixLength          = 2
	suffixLength          = 2
)

var defaultDelimiters = []string{",", ":"}
