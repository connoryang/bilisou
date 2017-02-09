package model

import (
	"fmt"
)

type UserVar struct {
	UK   int64
	Type int
	Page int
	ShareVars []ShareVar
}

func GetUserVar(Type int, Keyword string, Page int) *UserVar{
	fmt.Println("search share var called")
	return nil
}
