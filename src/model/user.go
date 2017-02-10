package model

import (
	"fmt"
	u "utils"
	//	"github.com/siddontang/go/log"
	"database/sql"
	"math"
)

type UserPageVar struct {
	Page   int
	Start  int
	End    int
	Previous int
	Next   int
	Before []int
	After  []int

	User        User
	Shares      []Share
	RandomUsers  []User
	RandomShares []Share
}

func GenerateUserPageVar(db *sql.DB, uk string, p int) *PageVar{
	pv := PageVar{}
	pv.Type = "user"

	s := "select uk, uname, avatar_url, fans_count, follow_count, pubshare_count, intro from uinfo where uk = %s"
	s = fmt.Sprintf(s, uk)


	users := GetUserBySql(db, s)
	if(len(users) != 1) {
		return nil
	}
	pv.User = users[0]

	pv.Current = p
	s = "SELECT count(s.id) FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where u.uk = %s"
	s = fmt.Sprintf(s, uk)
	found := GetFound(db, s)
	//	log.Error(found)

	d := float64(found) / float64(u.PAGEMAX)
	pv.End = int(math.Ceil(d))

	pv.Previous = pv.Current - 1;
	pv.Next = pv.Current + 1;


	SetBA(&pv)

	s = "SELECT s.data_id, s.title, s.share_id, s.album_id, u.uname, s.category, s.file_count, s.filenames, s.size, s.feed_time, s.view_count, s.like_count, s.last_scan, u.uk FROM sharedata as s join uinfo as u on s.uinfo_id = u.id where u.uk = %s order by last_scan desc limit %d, %d"
	s = fmt.Sprintf(s, uk, 0, 20)

	shares := GetShareBySql(db, s)
	pv.UserShares = shares

	pv.RandomShares = GenerateRandomShares(db, 10, 0, "")
	pv.RandomUsers = GenerateRandomUsers(db, 24)

	return &pv

}
