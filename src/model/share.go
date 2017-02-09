package model

import (
	"fmt"
	sql "database/sql"
	"github.com/siddontang/go/log"
	u "utils"
)

type ShareVar struct {
	ShareID   string
	DataID    string
	Title     string
	FeedType  string //专辑：album 文件或者文件夹：share
	AlbumID   string
	Category  string
	FeedTime  string
	Size      string
	Filenames string
	UK        string
	Uname     string
}

func GetShareVar(db *sql.DB, did string) *ShareVar{
	fmt.Println("get share var called")
	sql := "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.data_id = '" + did + "'"
	log.Info(sql)
	rows, err := db.Query(sql)
	u.CheckErr(err)
	sv := ShareVar{}
	for rows.Next() {
		var dataid sql.NullString
		var title sql.NullString
		var shareid sql.NullString
		var albumid sql.NullString
		var uname sql.NullString
		var filenames sql.NullString

		var category sql.NullInt64
		var fc sql.NullInt64
		var size sql.NullInt64
		var feedtime sql.NullInt64
		var vc sql.NullInt64
		var lc sql.NullInt64
		var ls sql.NullInt64

		err = rows.Scan(&dataid, &title, &shareid, &albumid, &uname, &category, &fc, &filenames, &size, &feedtime, &vc, &lc, &ls)

		u.CheckErr(err)

		if dataid.Valid {
			sv.DataID = dataid.String
		}

		fmt.Println(sv)
		fmt.Println(fc)
		fmt.Println(feedtime)
		//fmt.Println(avatarurl)
	}

	return nil
}
