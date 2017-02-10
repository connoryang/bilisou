package model

import (
	"fmt"
	"database/sql"
	"github.com/siddontang/go/log"
	u "utils"
	"math/rand"
	"time"
)

type Share struct {
	ShareID   string
	DataID    string
	Title     string
	FeedType  string //专辑：album 文件或者文件夹：share
	AlbumID   string
	Category  string
	CategoryInt int
	CategoryCN  string
	FeedTime  string
	Size      string
	Filenames []string
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
	var uk sql.NullInt64

	log.Info(s)
	rows, err := db.Query(s)
	u.CheckErr(err)

	shares := []Share{}

	for rows.Next() {
		sv := Share{}
		err = rows.Scan(&dataid, &title, &shareid, &albumid, &uname, &category, &fc, &filenames, &size, &feedtime, &vc, &lc, &ls, &uk)

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
			sv.CategoryInt = int(category.Int64)
		} else {
			sv.CategoryInt = 0
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


		if filenames.Valid {
			sv.Filenames = u.SplitNames(filenames.String)
		} else {
			sv.Filenames = []string{}
		}


		if fc.Valid {
			sv.FileCount = u.IntToStr(fc.Int64)
		} else {
			sv.FileCount = "0"
		}

		if size.Valid {
			sv.Size = u.SizeToStr(size.Int64)
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

		if uk.Valid {
			sv.UK = u.IntToStr(uk.Int64)
		} else {
			sv.UK = "0"
		}
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

func GetUserBySql(db *sql.DB, s string) ([]User){
	var UK          sql.NullInt64
	var Uname       sql.NullString
	var FansCount   sql.NullInt64
	var FollowCount sql.NullInt64
	var PubshareCount    sql.NullInt64
	var AvatarURL   sql.NullString
	var Intro       sql.NullString


	log.Info(s)
	rows, err := db.Query(s)
	u.CheckErr(err)

	users := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&UK, &Uname, &AvatarURL, &FansCount, &FollowCount, &PubshareCount, &Intro)

		if Uname.Valid {
			user.Uname = Uname.String
		} else {
			user.Uname = ""
		}

		if AvatarURL.Valid {
			user.AvatarURL = AvatarURL.String
		} else {
			user.AvatarURL = ""
		}

		if Intro.Valid {
			user.Intro = Intro.String
		} else {
			user.Intro = ""
		}

		if UK.Valid {
			user.UK = u.IntToStr(UK.Int64)
		} else {
			user.UK = "0"
		}

		if FollowCount.Valid {
			user.FollowCount = u.IntToStr(FollowCount.Int64)
		} else {
			user.FollowCount = "0"
		}

		if FansCount.Valid {
			user.FansCount = u.IntToStr(FansCount.Int64)
		} else {
			user.FansCount = "0"
		}

		if PubshareCount.Valid {
			user.PubshareCount = u.IntToStr(PubshareCount.Int64)
		} else {
			user.PubshareCount = "0"
		}
		users = append(users, user)

	}

	return users
}

func GetShareMaxMinID(db *sql.DB) (int, int) {
	var max int
	var min int
	sql := "select max(id), min(id) from sharedata"
	rows, err := db.Query(sql)
	u.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&max, &min)
	}
	return max, min
}

func GetUserMaxMINID(db *sql.DB) (int, int) {
	var max int
	var min int
	sql := "select max(id), min(id) from uinfo"
	rows, err := db.Query(sql)
	u.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&max, &min)
	}
	return max, min
}

func GenerateRandomShares(db *sql.DB, size int, category int, keyword string) []Share {
	max, min := GetShareMaxMinID(db)
	rs := []string{}
	for i := 0; i < size; i ++ {
		rand.Seed(time.Now().UnixNano())
		r := u.IntToStr(int64(rand.Intn(max - min) + min))
		rs = append(rs, r)
	}

	s := "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where s.id in ("
	for v, r := range rs {
		if v == (len(rs) - 1) {
			s = s + r
		} else {
			s = s + r + ", "
		}
	}
	s = s + ")"
	shares := GetShareBySql(db, s)
	return shares
}


func GenerateRandomUsers(db *sql.DB, size int) []User {
	max, min := GetUserMaxMINID(db)
	rs := []string{}
	for i := 0; i < size; i ++ {
		rand.Seed(time.Now().UnixNano())
		r := u.IntToStr(int64(rand.Intn(max - min) + min))
		rs = append(rs, r)
	}

	s := "select uk, uname, avatar_url, fans_count, follow_count, pubshare_count, intro from uinfo where id in ("
	for v, r := range rs {
		if v == (len(rs) - 1) {
			s = s + r
		} else {
			s = s + r + ", "
		}
	}
	s = s + ")"
	users := GetUserBySql(db, s)
	return users
}


func UpdateCategory(db *sql.DB) {
	max, min := GetShareMaxMinID(db)
	for i:=min; i <= max; i ++ {
		s := "select title from sharedata where id = %d"

		s = fmt.Sprintf(s, i)
		rows, err := db.Query(s)

		u.CheckErr(err)
		var tt sql.NullString

		for rows.Next() {
			err = rows.Scan(&tt)
		}
		if tt.Valid {
			c := u.GetCategoryFromName(tt.String)
			us := "update sharedata set category = ? where id = ?"
			//us = fmt.Sprintf(us, c, i)
			//db.Query(us)
			stmt, _ := db.Prepare(us)
			stmt.Exec(c,i)
			stmt.Close()
			//res.RowsAffected()

			//db.Exec(us)
			log.Info(us)
		}
	}
	log.Info("ajl")
}
