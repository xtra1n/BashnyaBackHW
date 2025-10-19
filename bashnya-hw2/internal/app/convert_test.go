package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ConvertNumToTextTest(t *testing.T) {
	testTable := []struct {
		num      int
		excepted string
	}{
		{0, "Ноль"},
		{1, "один"},
		{11, "одинадцать"},
		{23, "двадцать три"},
		{405, "четыреста пять"},
		{2369, "две тысячи триста шестьдесят девять"},
		{81893, "восемьдесят одна тысяча восемьсот девяносто три"},
		{500000, "пятьсот тысяч"},
	}

	for _, testCase := range testTable {
		result := ConvertNumToText(testCase.num)
		assert.Equal(t, testCase.excepted, result)
	}
}
