package model


import (
//	"fmt"
//	"database/sql"
//	"github.com/siddontang/go/log"
	u "utils"
//	"math/rand"
//	"time"
)


type PageVar struct {
	Type          string
	CategoryInt   int
	Category      string
	CategoryCN    string

	//for paging
	Current       int
	Start         int
	End    int
	Previous int
	Next   int
	Before []int
	After  []int


	User              User
	Share             Share
	ListShares        []Share
	ListUsers         []User
	RandomUsers       []User
	UserShares        []Share
	RandomShares      []Share
	RandomSharesCategory      []Share
	RandomSharesSimilar      []Share

}

func SetBA(pv *PageVar) {
	pp := pv.Current - u.NAVMAX
	for ; (pv.Current > pp) && (pp >= 1); pp ++ {
		pv.Before = append(pv.Before, pp)
	}

	pp = pv.Current + 1
	for ; ((pv.Current + u.NAVMAX) >= pp) && (pp <= pv.End); pp ++ {
		pv.After = append(pv.After, pp)
	}

}
