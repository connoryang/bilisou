package model

import (
	"fmt"
)

type SearchPageVar struct {
	Type int
	Keyword string
	Page int
	Shares []Share
}

func GenerateSearchPageVar(Type int, Keyword string, Page int) *SearchPageVar {
	fmt.Println("search var called")
	return nil
}
