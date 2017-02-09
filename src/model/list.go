package model

import (
	"fmt"
	u "utils"
	"github.com/siddontang/go/log"
	"database/sql"
	"math"
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
	Users  []User
	Before []int
	After  []int
}

func GenerateListPageVar(db *sql.DB, c int, p int) *ListPageVar{

	if p <= 0 {
		return nil
	}
	var sql string
	var sqlfound string

	lp := ListPageVar{}

	lp.CategoryInt = c

	cat, ok := u.CAT_INT_STR[c]
	if ok {
		lp.Category = cat
	}

	cat, ok = u.CAT_INT_STRCN[c]
	if ok {
		lp.CategoryCN = cat
	}


	lp.Page = p
	lp.Start = 1

	if c == 0 {
		sql = "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan FROM sharedata as s join uinfo as u on s.uinfo_id = u.id order by last_scan desc limit %d, %d"
		sql = fmt.Sprintf(sql, (lp.Page - 1) * u.PAGEMAX, u.PAGEMAX)
		sqlfound = "SELECT count(s.id)  FROM sharedata as s join uinfo as u on s.uinfo_id = u.id order by last_scan desc"
		sqlfound = fmt.Sprintf(sqlfound)
	} else {
		sql = "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.category = %d order by last_scan desc limit %d, %d"
		sql = fmt.Sprintf(sql, lp.CategoryInt, (lp.Page - 1) * u.PAGEMAX, u.PAGEMAX)
		sqlfound = "SELECT count(s.id) FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.category = %d order by last_scan desc"
		sqlfound = fmt.Sprintf(sqlfound, lp.CategoryInt)
	}


	shares := GetShareBySql(db, sql)
	found := GetFound(db, sqlfound)
	log.Error(found)

	d := float64(found) / float64(u.PAGEMAX)
	lp.End = int(math.Ceil(d))
	lp.Previous = lp.Page - 1;
	lp.Next = lp.Page + 1;


	for _, s := range shares {
		lp.Shares = append(lp.Shares, s)
	}

	pp := lp.Page - u.NAVMAX
	for ; (lp.Page > pp) && (pp >= 1); pp ++ {
		lp.Before = append(lp.Before, pp)
	}


	pp = lp.Page + 1
	for ; ((lp.Page + u.NAVMAX) >= pp) && (pp <= lp.End); pp ++ {
		lp.After = append(lp.After, pp)
	}


	return &lp
}
