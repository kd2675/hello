package act

import (
	"calculator/biz/parser"
	"calculator/biz/service"
	"calculator/biz/validator"
)

type AddCalculator struct {
	delimiterParser  parser.DelimiterParser
	numberParser     parser.NumberParser
	parsingValidator validator.ParsingValidator
	sumService       service.SumService
}

func NewAddCalculator(
	delimiterParser parser.DelimiterParser,
	numberParser parser.NumberParser,
	parsingValidator validator.ParsingValidator,
	sumService service.SumService,
) *AddCalculator {
	return &AddCalculator{
		delimiterParser:  delimiterParser,
		numberParser:     numberParser,
		parsingValidator: parsingValidator,
		sumService:       sumService,
	}
}

func (calculator *AddCalculator) Add(input string) (int64, error) {
	if input == "" {
		return 0, nil
	}

	delimiters, _ := calculator.delimiterParser.ParseDelimiters(input)

	body, _ := calculator.delimiterParser.ExtractBody(input)

	numbers, err := calculator.numberParser.ParseNumbers(body, delimiters)
	if err != nil {
		return 0, err
	}

	err = calculator.parsingValidator.Validate(numbers)
	if err != nil {
		return 0, err
	}

	result := calculator.sumService.Sum(numbers)

	return result, nil
}
