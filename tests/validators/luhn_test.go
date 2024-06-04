package validators_test 

import(
	"testing"
	"github.com/fohalloran/go-vard-validator/internal/validators"
)

func TestValidNumbers (t *testing.T){
	validNumbers := []int{4242424242424242,4000056655665556,5555555555554444,2223003122003222,5200828282828210,5105105105105100,6011111111111117,6011000990139424,6011981111111113,3056930009020004,36227206271667,6555900000604105,3566002020360505,6200000000000005,6200000000000047} 
	for _, element := range validNumbers {
		valid := validators.LuhnValidation(element)
		if valid != true{
			t.Errorf("Error: %d should be valid", element)
		}
	}
	
}
func TestInvalidNumbers (t *testing.T){
	validNumbers := []int{4242424242424243,4000056655665557,5555555555554445,2223003122003223,5200828282828211,5105105105105101,6011111111111118,6011000990139425,6011981111111114,3056930009020005,36227206271668,6555900000604106,3566002020360506,6200000000000006,6200000000000048} 
	for _, element := range validNumbers {
		valid := validators.LuhnValidation(element)
		if valid != false{
			t.Errorf("Error: %d should be invalid", element)
		}
	}
	
}