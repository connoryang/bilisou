package model

import (
	"fmt"
	"database/sql"
	"github.com/siddontang/go/log"
	u "utils"
)

type Share struct {
	ShareID   string
	DataID    string
	Title     string
	FeedType  string //专辑：album 文件或者文件夹：share
	AlbumID   string
	Category  string
	CategoryCN  string
	FeedTime  string
	Size      string
	Filenames string
	FileCount string
	UK        string
	Uname     string
	ViewCount string
	LikeCount string
	LastScan  string
}

func GetShareBySql(db *sql.DB, s string) ([]Share){
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

	log.Info(s)
	rows, err := db.Query(s)
	u.CheckErr(err)

	shares := []Share{}

	for rows.Next() {
		sv := Share{}
		err = rows.Scan(&dataid, &title, &shareid, &albumid, &uname, &category, &fc, &filenames, &size, &feedtime, &vc, &lc, &ls)

		u.CheckErr(err)

		if dataid.Valid {
			sv.DataID = dataid.String
		} else {
			sv.DataID = ""
		}

		if title.Valid {
			sv.Title = title.String
		} else {
			sv.Title = ""
		}

		if shareid.Valid {
			sv.ShareID = shareid.String
		} else {
			sv.ShareID = ""
		}

		if albumid.Valid {
			sv.AlbumID = albumid.String
		} else {
			sv.AlbumID = ""
		}

		if uname.Valid {
			sv.Uname = uname.String
		} else {
			sv.Uname = ""
		}

		if category.Valid {
			//sv.Category = u.IntToStr(category.Int64)
			cat, ok := u.CAT_INT_STR[int(category.Int64)]
			if ok {
				sv.Category = cat
			}
		} else {
			sv.Category = ""
		}


		if category.Valid {
			//sv.Category = u.IntToStr(category.Int64)
			cat, ok := u.CAT_INT_STRCN[int(category.Int64)]
			if ok {
				sv.CategoryCN = cat
			}
		} else {
			sv.Category = ""
		}

		if fc.Valid {
			sv.FileCount = u.IntToStr(fc.Int64)
		} else {
			sv.FileCount = "0"
		}

		if size.Valid {
			sv.Size = u.IntToStr(size.Int64)
		} else {
			sv.Size = "0"
		}

		if vc.Valid {
			sv.ViewCount = u.IntToStr(vc.Int64)
		} else {
			sv.ViewCount = "0"
		}


		if lc.Valid {
			sv.LikeCount = u.IntToStr(lc.Int64)
		} else {
			sv.LikeCount = "0"
		}


		if feedtime.Valid {
			sv.FeedTime = u.IntToDateStr(feedtime.Int64)
		} else {
			sv.FeedTime = "0"
		}

		if ls.Valid {
			sv.LastScan = u.IntToDateStr(ls.Int64)
		} else {
			sv.LastScan = "0"
		}

		fmt.Println(sv)
		shares = append(shares, sv)
	}
	return shares
}

func GetFound(db *sql.DB, sql string) int {
	var found int
	log.Info(sql)
	rows, err := db.Query(sql)
	u.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&found)
	}
	return found
}

type User struct {
	UK          string
	Uname       string
	FansCount   string
	FollowCount string
	PubshareCount    string
	AvatarURL   string
	Intro       string
}

func GetUserBySql(category int, Keyword string, Page int) *User{
	fmt.Println("search share var called")
	return nil
}
