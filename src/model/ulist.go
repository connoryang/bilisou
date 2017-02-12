package model

import (
//	"fmt"
//	"database/sql"
	es "gopkg.in/olivere/elastic.v3"
//	"github.com/siddontang/go/log"
//	u "utils"
//	"math"
//	"time"
)


func GenerateUlistPageVar(esclient *es.Client, page int) *PageVar {
	if page <= 0 {
		return nil
	}

	pv := PageVar{}
	pv.Type = "ulist"


	query := es.NewTermQuery("search", 1)
	var size int64
	pv.ListUsers, size = SearchUser(esclient, query, page, 20)

	if len(pv.ListUsers) == 0 {
		//return nil
		pv.Type = "lost"
	}

	//pv.End = int(size)
	pv.End = int(size) / 20 + 1
	pv.Current = page
	SetBA(&pv)

	pv.RandomShares = GenerateRandomShares(esclient, 0, 10, "")
	pv.RandomUsers = GenerateRandomUsers(esclient, 24)
	pv.Keywords = GenerateRandomKeywords(esclient, 30)
	return &pv
}


/*

func GenerateUlistPageVar1(db *sql.DB, p int) *PageVar {
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
*/
