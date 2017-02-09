package model

import (
	"fmt"
)


type ListVar struct {
	Type  int
	Page  int
	ShareVars []ShareVar

}

func GetListVar(Type int, Page int) *ListVar{
	fmt.Println("get list var called")
	return nil
}
