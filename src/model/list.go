package model

import (
//	"fmt"
//	u "utils"
//	"github.com/siddontang/go/log"
//	"database/sql"
	es "gopkg.in/olivere/elastic.v3"
//	"math"
)



func GenerateListPageVar(esclient *es.Client, category int,  page int) *PageVar {
	pv := PageVar{}
	if page <= 0 {
		pv.Type = "lost"
		pv.RandomShares = GenerateRandomShares(esclient, 0, 10, "")
		pv.RandomUsers = GenerateRandomUsers(esclient, 24)
		pv.Keywords = GenerateRandomKeywords(esclient, 30)

		return &pv
	}


	pv.Type = "list"

	boolQuery := es.NewBoolQuery()
	query := es.NewTermQuery("search", 1)
	boolQuery.Should(query)
	if category != 0 {
		boolQuery.Must(es.NewTermQuery("category", category))
	}
	var size int64
	pv.ListShares, size = SearchShare(esclient, boolQuery, page, 20, "last_scan")

	if len(pv.ListShares) == 0 {
		//return nil
		pv.Type = "lost"
	}

	pv.End = int(size) / 20 + 1
	pv.Current = page

	SetBA(&pv)
	SetCategory(&pv, category)


	pv.RandomShares = GenerateRandomShares(esclient, 0, 10, "")
	pv.RandomUsers = GenerateRandomUsers(esclient, 24)
	pv.Keywords = GenerateRandomKeywords(esclient, 30)
	return &pv
}


/*
func GenerateListPageVar1(db *sql.DB, c int, p int) *PageVar{

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
*/
