package nhsnum

import (
	"strconv"
	"strings"
)

func IsValid(num string) bool {
	digits := strings.Split(num, "")
	length := len(digits)

	total := calculateTotal(digits)
	providedCheckNum := strDigitToInt(digits[length-1])
	remainder := total % 11
	calcedCheckNum := 11 - remainder

	// providedCheckNum is always going to be a single digit, no need to explicitly
	// check if calcedCheckNum is 10. The default will catch that scenario
	if providedCheckNum == calcedCheckNum {
		return true
	}

	// 11 is a magic number but I'm not sure what the terminology for it is or I'd create a variable
	return providedCheckNum == 0 && calcedCheckNum == 11
}

func calculateTotal(nums []string) int {
	length := len(nums)
	total := 0

	for i := 0; i < length-1; i++ {
		num := strDigitToInt(nums[i])

		factor := length - i
		total += num * factor
	}

	return total
}

func strDigitToInt(digit string) int {
	num, err := strconv.Atoi(digit)

	if err != nil {
		panic(err)
	}

	return num
}
