package model

import (
//	"fmt"
//	"database/sql"
//	"github.com/siddontang/go/log"
	u "utils"
	t "html/template"
)

type ShareData struct {
	Id         int64
	Album_id   int64
	Category   int64
	Data_id  string
	Feedtime int64
	File_count int64
	Filenames string
	Last_scan int64
	Like_count int64
	Share_id   string
	Size       int64
	Title     string
	HTitle    string
	Uinfo_id   int64
	Uk         string
	Uname     string
	View_count int64
}

type UserInfo struct {
	Id           int64
	Avatar_url   string
	Fans_count   int64
	Follow_count int64
	Intro        string
	Pubshare_count  int64
	Uk            string
	Uname         string
}

type Share struct {
	ShareID   string
	HTitle    t.HTML
	DataID    string
	Title     string
	FeedType  string //专辑：album 文件或者文件夹：share
	AlbumID   string
	Category  string
	CategoryInt int64
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

type User struct {
	UK          string
	Uname       string
	FansCount   string
	FollowCount string
	PubshareCount    string
	AvatarURL   string
	Intro       string
}

type Keyword struct {
	Keyword string
	Count   int64
}

func ShareDataToShare(sd ShareData) Share {
/*	Id         int64
	Album_id string
	Category string
	Data_id  string
	Feedtime int64
	File_count int64
	Filenames string
	Last_scan int64
	Like_count int64
	Share_id   string
	size       int64
	title     string
	Uinfo_id   int64
	Uk         string
	Uname     string
	View_count int64
*/
	s := Share{}
	s.ShareID = sd.Share_id
	s.DataID  = sd.Data_id
	s.Title    = sd.Title
	s.HTitle = t.HTML(sd.HTitle)

	s.AlbumID  = u.IntToStr(sd.Album_id)
	s.CategoryInt  = sd.Category
	s.Category = u.CAT_INT_STR[int(sd.Category)]
	s.CategoryCN = u.CAT_INT_STRCN[int(sd.Category)]
	s.FeedTime  = u.IntToDateStr(sd.Feedtime)
	s.Size      = u.SizeToStr(sd.Size)
	s.Filenames = u.SplitNames(sd.Filenames)
	s.FileCount = u.IntToStr(sd.File_count)
	s.UK        = sd.Uk
	s.Uname     = sd.Uname
	s.ViewCount = u.IntToStr(sd.View_count)
	s.LikeCount   = u.IntToStr(sd.Like_count)
	s.LastScan    = u.IntToDateStr(sd.Last_scan)
	return s
}

func UserInfoToUser(uinfo UserInfo) User {
/*	Id           int64
	Avatar_url   string
	Fans_count   int64
	Follow_count int64
	Intro        string
	Pubshare_count  int64
	Uk            int64
	Uname         string
*/
	user := User{}
	user.UK         =  uinfo.Uk
	user.Uname      =  uinfo.Uname
	user.FansCount  =  u.IntToStr(uinfo.Fans_count)
	user.FollowCount = u.IntToStr(uinfo.Follow_count)
	user.PubshareCount   = u.IntToStr(uinfo.Pubshare_count)
	user.AvatarURL   =  uinfo.Avatar_url
	user.Intro       =  uinfo.Intro
	return user

}

func SetCategory(pv *PageVar, category int){
	pv.CategoryInt = category

	cat, ok := u.CAT_INT_STR[category]
	if ok {
		pv.Category = cat
	}

	cat, ok = u.CAT_INT_STRCN[category]
	if ok {
		pv.CategoryCN = cat
	}
}
