package roman

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type Roman struct {
	value uint16
}

func NewRoman(arabic uint16) (Roman, error) {
	if arabic < 1 || arabic > 3999 {
		return Roman{}, fmt.Errorf("cannot create a roman numeral for numbers outside the range 1-3999")
	}
	return Roman{value: arabic}, nil
}

func (r Roman) String() string {
	var result strings.Builder
	arabic := r.value
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func (r Roman) Uint16() uint16 {
	return r.value
}

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", fmt.Errorf("value %d is too large, must be less than 4000", arabic)
	}

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}