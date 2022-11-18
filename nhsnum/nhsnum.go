package nhsnum

import (
	"strconv"
	"strings"
)

// first pass logic, todo: refactor
func IsValid(num string) bool {
	digits := strings.Split(num, "")
	length := len(digits)

	total := 0
	for i := 0; i < length-1; i++ {
		num, err := strconv.Atoi(digits[i])

		if err != nil {
			panic(err)
		}

		factor := length - i
		total += num * factor
	}

	remainder := total % 11
	providedCheckNum, _ := strconv.Atoi(digits[length-1])
	calcedCheckNum := 11 - remainder

	if providedCheckNum == calcedCheckNum {
		return true
	}

	if providedCheckNum == 0 && calcedCheckNum == 11 {
		return true
	}

	return false
}
