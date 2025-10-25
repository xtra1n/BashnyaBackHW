package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessNumber(t *testing.T) {
	testTable := []struct {
		num               int
		excepted          int
		exceptError       bool
		exceptedErrorType string
	}{
		{-12306, 12307, false, ""},
		{12299, 479662, false, ""},
		{12303, 159940, false, ""},
		{12305, 36922, false, ""},
		{63, -0, true, "service error"},
	}

	for _, testCase := range testTable {
		result, err := ProcessNumber(testCase.num)

		if testCase.exceptError == false {
			assert.Equal(t, testCase.excepted, result)
		} else {
			assert.EqualError(t, err, testCase.exceptedErrorType)
		}
	}
}
