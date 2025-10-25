package app

import (
	"strings"
)

var (
	unitsMale = []string{
		"", "один", "два", "три", "четыре", "пять",
		"шесть", "семь", "восемь", "девять",
	}

	unitsFemale = []string{
		"", "одна", "две", "три", "четыре", "пять",
		"шесть", "семь", "восемь", "девять",
	}

	teens = []string{
		"десять", "одиннадцать", "двенадцать", "тринадцать",
		"четырнадцать", "пятнадцать", "шестнадцать",
		"семнадцать", "восемнадцать", "девятнадцать",
	}

	tens = []string{
		"", "", "двадцать", "тридцать", "сорок", "пятьдесят",
		"шестьдесят", "семьдесят", "восемьдесят", "девяносто",
	}

	hundreds = []string{
		"", "сто", "двести", "триста", "четыреста",
		"пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот",
	}
)

func ConvertNumToText(num int) string {
	if num == 0 {
		return "Ноль"
	}

	var text strings.Builder

	thousands := num / 1000

	if thousands > 0 {
		writeHundreds(&text, thousands, true)
		text.WriteString(" ")
		text.WriteString(selectThousandsCase(thousands))
		num %= 1000
		if num > 0 {
			text.WriteString(" ")
		}
	}

	if num > 0 {
		writeHundreds(&text, num, false)
	}

	return text.String()
}

func writeHundreds(text *strings.Builder, num int, isFemale bool) {
	if num >= 100 {
		text.WriteString(hundreds[num/100])
		num %= 100
		if num > 0 {
			text.WriteString(" ")
		}
	}

	if num >= 20 {
		text.WriteString(tens[num/10])
		num %= 10
		if num > 0 {
			text.WriteString(" ")
		}
	}

	if num >= 10 {
		text.WriteString(teens[num-10])
		return
	}

	if num > 0 {
		switch isFemale {
		case true:
			text.WriteString(unitsFemale[num])
		case false:
			text.WriteString(unitsMale[num])
		}
	}
}

func selectThousandsCase(num int) string {
	LastTwoNums := num % 100
	LastNum := num % 10

	var result string

	if LastTwoNums == 11 || LastTwoNums == 14 {
		result = "тысяч"
	}

	if result == "" {
		switch LastNum {
		case 1:
			result = "тысяча"
		case 2, 3, 4:
			result = "тысячи"
		default:
			result = "тысяч"
		}
	}

	return result
}
