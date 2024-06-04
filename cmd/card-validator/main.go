package main

import (
	"fmt"
	"github.com/fohalloran/go-vard-validator/internal/validators"
)

func main(){
	valid := validators.LuhnValidation(32409857)
	fmt.Println(valid)
}