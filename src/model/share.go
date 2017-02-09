package model

import (
	"fmt"
)

type SharePageVar struct {
	DataID      string
	Shares      []Share
	Users       []User
}

func GenerateShareVar(Type int, Keyword string, Page int) *SharePageVar{
	fmt.Println("page share var called")
	return nil
}
