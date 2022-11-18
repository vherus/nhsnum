package nhsnum

import (
	"strconv"
	"strings"
)

// first pass logic, todo: refactor
func IsValid(num string) bool {
	digits := strings.Split(num, "")
	length := len(digits)

	total := calculateTotal(digits)

	remainder := total % 11
	providedCheckNum, _ := strconv.Atoi(digits[length-1])
	calcedCheckNum := 11 - remainder

	// providedCheckNum is always going to be a single digit, no need to explicitly
	// check if calcedCheckNum is 10. The default will catch that scenario
	if providedCheckNum == calcedCheckNum {
		return true
	}

	if providedCheckNum == 0 && calcedCheckNum == 11 {
		return true
	}

	return false
}

func calculateTotal(nums []string) int {
	length := len(nums)
	total := 0

	for i := 0; i < length-1; i++ {
		num, err := strconv.Atoi(nums[i])

		if err != nil {
			panic(err)
		}

		factor := length - i
		total += num * factor
	}

	return total
}
