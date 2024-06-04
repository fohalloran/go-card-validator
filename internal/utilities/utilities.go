package utilities

import "strconv"

func SplitIntToSlice(number int) []int {

	if number < 0 {
		number = -number
	}

	length := len(strconv.Itoa(number))
	output := make([]int, length)
	for i := 0; i < length; i++ {
		output[length-i-1] = number % 10 // Populate the array from the end to the start with the number on the end
		number = number / 10
	}
	return output

}
