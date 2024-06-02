package main

import (
	"strconv"
)

func splitIntToSlice(number int, length int)[]int{

	output := make([]int, length)
	for i:=0;i < length;i++{
		output[length-i-1] = number % 10 // Populate the array from the end to the start with the number on the end
		number = number / 10
	}
	return output

}

func LuhnValidation(cardNumber int)bool{
	sum := 0
	length := len(strconv.Itoa(cardNumber))
	cardSlice := splitIntToSlice(cardNumber, length)
	
	for i:= length-1;i>=0;i-- {
		number := cardSlice[i]
		if i % 2 == 0 {
			number *= 2
			if number > 9{
				number = number - 9
	
			}
		} 
		sum = sum + number
		
	}
	return sum % 10 == 0
}

func main(){
	
	
}

