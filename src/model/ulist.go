package model

import (
	"fmt"
	"database/sql"
//	"github.com/siddontang/go/log"
	u "utils"
	"math"
//	"time"
)


func GenerateUlistPageVar(db *sql.DB, p int) *PageVar {
	pv := PageVar{}
	pv.Type = "ulist"


	s := "select uk, uname, avatar_url, fans_count, follow_count, pubshare_count, intro from uinfo limit %s, %s"
	s = fmt.Sprintf(s, (pv.Current - 1) * u.PAGEMAX, u.PAGEMAX)


	users := GetUserBySql(db, s)
	if(len(users) != 1) {
		return nil
	}
	pv.ListUsers = users

	pv.Current = p
	s = "SELECT count(id) FROM uinfo"
	found := GetFound(db, s)

	d := float64(found) / float64(u.PAGEMAX)
	pv.End = int(math.Ceil(d))

	pv.Previous = pv.Current - 1;
	pv.Next = pv.Current + 1;

	SetBA(&pv)

	pv.RandomShares = GenerateRandomShares(db, 10, 0, "")
	pv.RandomUsers = GenerateRandomUsers(db, 24)

	return &pv

}
