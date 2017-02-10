package model

import (
	"fmt"
//	u "utils"
	//	"github.com/siddontang/go/log"
	"database/sql"

)

type SharePageVar struct {
	DataID      string
	Share       Share
	UserShares        []Share
	RandomSharesCategory      []Share
	RandomSharesSimilar      []Share
	RandomUsers       []User
}

func GenerateShareVar(db *sql.DB, dataid string) *PageVar{

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
