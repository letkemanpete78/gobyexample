package taskone

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// strNum := "10,\n200, -33,a.jjs,1001"
	// delimiter := ","
	// // var sum int64 = 0
	// if (strNum == "") || (delimiter == "") {

	// }
	// sum, err := SumValues(strNum, delimiter)
	// if err == nil {
	// 	fmt.Println(sum)
	// } else {
	// 	fmt.Println(err)
	// }
}

// SumValues - given numbers as a string, sum them up.
func SumValues(strNum string, delimiter string) (int64, error) {
	var sum int64
	numbers := strings.Split(strNum, delimiter)
	for num := range numbers {
		i, err := strconv.ParseInt(strings.Trim(numbers[num], "\n "), 10, 32)
		if err == nil {
			if i < 0 {
				return 0, fmt.Errorf("negative number %d cannot be used", i)
			} else if i < 1000 {
				sum = sum + i
			}
		}
	}
	return sum, nil
}
