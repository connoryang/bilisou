package model

import (
	"fmt"
)

type SearchVar struct {
	Type int
	Keyword string
	Page int
	ShareVars []ShareVar
}

func GetSearchVar(Type int, Keyword string, Page int) *SearchVar {
	fmt.Println("search var called")
	return nil
}
