package user

import "strings"

type User struct{
	Name string 
	Age int
	Address string
	Status bool
}

func NewUser(name string, age int, address string, status bool) User{
	return User{
		Name: name,
		Age: age,
		Address: address,
		Status: status,
	}
}
 
func (u User) CapitalizeUserName() string{
	return strings.ToUpper(u.Name)
}