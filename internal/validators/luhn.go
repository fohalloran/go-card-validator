package validators

import (
	"strconv"

	"github.com/fohalloran/go-vard-validator/internal/utilities"
)

func LuhnValidation(cardNumber int) bool {
	sum := 0
	length := len(strconv.Itoa(cardNumber))
	cardSlice := utilities.SplitIntToSlice(cardNumber)

	for i := length - 1; i >= 0; i-- {
		number := cardSlice[i]
		if i % 2 == 0 {
			number *= 2
			if number > 9 {
				number = number - 9

			}
		}
		sum = sum + number

	}
	return sum%10 == 0
}
