package model

import (
	"fmt"
	u "utils"
//	"github.com/siddontang/go/log"
	"database/sql"
//	"math"
)


type ListPageVar struct {
	CategoryInt   int
	Category string
	CategoryCN string
	Page   int
	Start  int
	End    int
	Previous int
	Next   int
	Shares []Share
	RandomUsers  []User
	Before []int
	After  []int
}

func GenerateListPageVar(db *sql.DB, c int, p int) *PageVar{

	if p <= 0 {
		return nil
	}
	var sql string
//	var sqlfound string

	pv := PageVar{}
	pv.Type = "list"

	pv.CategoryInt = c

	cat, ok := u.CAT_INT_STR[c]
	if ok {
		pv.Category = cat
	}

	cat, ok = u.CAT_INT_STRCN[c]
	if ok {
		pv.CategoryCN = cat
	}


	pv.Current = p
	pv.Start = 1

	if c == 0 {
		sql = "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id order by last_scan desc limit %d, %d"
		sql = fmt.Sprintf(sql, (pv.Current - 1) * u.PAGEMAX, u.PAGEMAX)
//		sqlfound = "SELECT count(s.id)  FROM sharedata as s join uinfo as u on s.uinfo_id = u.id order by last_scan desc"
//		sqlfound = fmt.Sprintf(sqlfound)
	} else {
		sql = "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.category = %d order by last_scan desc limit %d, %d"
		sql = fmt.Sprintf(sql, pv.CategoryInt, (pv.Current - 1) * u.PAGEMAX, u.PAGEMAX)
//		sqlfound = "SELECT count(s.id) FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.category = %d order by last_scan desc"
//		sqlfound = fmt.Sprintf(sqlfound, lp.CategoryInt)
	}


	shares := GetShareBySql(db, sql)
	pv.ListShares = shares

//	found := GetFound(db, sqlfound)
//	log.Error(found)

//	d := float64(found) / float64(u.PAGEMAX)
//	lp.End = int(math.Ceil(d))
	pv.End = 50
	pv.Previous = pv.Current - 1;
	pv.Next = pv.Current + 1;
	SetBA(&pv)


	pv.RandomShares = GenerateRandomShares(db, 10, 0, "")
	pv.RandomUsers = GenerateRandomUsers(db, 24)

	return &pv
}
