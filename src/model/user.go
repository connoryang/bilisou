package model

import (
	"fmt"
)

type UserPageVar struct {
	Type        string
	Page        string
	Shares      []Share
	Users       []User
}

func GenerateUserPageVar(Type int, Keyword string, Page int) *UserPageVar{
	fmt.Println("search share var called")
	return nil
}
