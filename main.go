package main

import (
	c_errors "class_proj_secb/customErrors"
	user "class_proj_secb/userModule"
	"errors"
	"log"
	"strconv"
)

func main() {
	println("Hello, World!")
	_, err := divideWithError(10,0)
	if err != nil{
		println(err.Error())
	}
	divideResult, _ := divideWithError(30,19)
	println(*divideResult)
	// importing user
	user := user.NewUser( "John", 20, "New York", true)
	log.Println("User : ", user)
	cName := user.CapitalizeUserName()
	println("CApitalized User Name : ", cName)
	// parse String
	_, err = parseString("abcd")
	println("error: ", err.Error())
	_, err = parseString("2000")
	println("error: ", err.Error())
	result, _ := parseString("100")
	println("result: ", *result)
}

func parseString(value string) (*int, error){
	res, err := strconv.Atoi(value)
	if err != nil{
		return nil, c_errors.CustomError{
			Message: "Cannot parse string to int",
			Type: "strconv Atoi",
			Err: err,
		}
	}
	if res < 1 || res > 1000{
		return nil, c_errors.CustomError{
			Message: "Value must be between 1 and 1000",
			Type: "range (1-1000)",
			Err: err,
		}
	}
	return &res, nil
}

func divideWithError(a, b float64) (*float64, error){
	if b == 0{
		return nil, errors.New("cannot divisible by zero")
	}
	res := a/b
	return &res, nil
}