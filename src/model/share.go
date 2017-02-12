package model

import (
//	"fmt"
//	u "utils"
		"github.com/siddontang/go/log"
//	"database/sql"
	es "gopkg.in/olivere/elastic.v3"

)


func GenerateSharePageVar(esclient *es.Client, dataid string) *PageVar {
	pv := PageVar{}
	pv.Type = "share"

	query := es.NewTermQuery("data_id", dataid)
	Shares, size := SearchShare(esclient, query, 1, 10, "")
	log.Info(Shares,"------\n", size)

	if len(Shares) == 0 {
		//return nil
		pv.Type = "lost"
	} else {
		pv.Share = Shares[0]
	}

	pv.RandomSharesSimilar = GenerateRandomShares(esclient, 0, 10, pv.Share.Title)
	pv.RandomSharesCategory = GenerateRandomShares(esclient, 0, 10, "")

	pv.RandomShares = GenerateRandomShares(esclient, 0, 10, "")
	pv.RandomUsers = GenerateRandomUsers(esclient, 24)
	pv.Keywords = GenerateRandomKeywords(esclient, 30)
	return &pv

}

/*
func GenerateShareVar1p(db *sql.DB, dataid string) *PageVar{

	sql := "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where data_id = %s"
	sql = fmt.Sprintf(sql, dataid)

	shares := GetShareBySql(db, sql)
	if len(shares) == 0 {
		return nil
	}

	pv := PageVar{}
	pv.Type = "share"

	pv.Share = shares[0]


	s := "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where u.uk = %s order by last_scan desc limit 0, 10"
	s = fmt.Sprintf(s, pv.Share.UK)

	pv.UserShares = GetShareBySql(db, s)

	pv.RandomSharesSimilar = GenerateRandomShares(db, 10, 0, pv.Share.Title)
	pv.RandomSharesCategory = GenerateRandomShares(db, 10, pv.Share.CategoryInt, "")

	pv.RandomShares = GenerateRandomShares(db, 10, 0, "")
	pv.RandomUsers = GenerateRandomUsers(db, 24)

	return &pv
}

*/
