package main

import (
	"calculator/act"
	"calculator/biz/parser"
	"calculator/biz/service"
	"calculator/biz/validator"
	"calculator/utils"
	"fmt"
)

func main() {
	input, err := utils.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	calc := act.NewAddCalculator(
		parser.NewDelimiterParserPattern(),
		parser.NewNumberParser(),
		validator.NewParsingValidator(),
		service.NewSumService(),
	)

	result, err := calc.Add(input)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	utils.PrintResult(result)
}
