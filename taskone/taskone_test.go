package taskone

import (
	"testing"
)

func TestSumValuesCorrect(t *testing.T) {
	strNum := "1,2,3"
	var expected int64 = 6
	delimiter := ","
	sum, err := SumValues(strNum, delimiter)
	if err == nil {
		if sum != expected {
			t.Errorf("sum error")
		}
	}
}

func TestSumValuesNewline(t *testing.T) {
	strNum := "1\n,-2,3"
	var expected int64 = 4
	delimiter := ","

	sum, err := SumValues(strNum, delimiter)
	if err == nil {
		if expected != sum {
			t.Errorf("newline error")
		}
	}

}

func TestSumValuesBigNum(t *testing.T) {
	strNum := "1\n,2000,3"
	var expected int64 = 4
	delimiter := ","

	sum, err := SumValues(strNum, delimiter)
	if err == nil {
		if expected != sum {
			t.Errorf("sum error [1000]")
		}
	}

}

func TestSumValuesNagative(t *testing.T) {
	strNum := "1,-2,3"
	var expected int64
	delimiter := ","

	sum, err := SumValues(strNum, delimiter)

	if expected != sum {
		t.Errorf("sum error")
	}

	if err == nil || err.Error() != "negative number -2 cannot be used" {
		t.Errorf("exception not thrown")
	}
}
