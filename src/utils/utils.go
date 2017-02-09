package utils

import (

	//"fmt"
	//	_ "github.com/go-sql-driver/mysql"
	//sql "database/sql"
	//	"io/ioutil"
	log "github.com/siddontang/go/log"
	//	"regexp"
	//	"encoding/json"
	//	"time"
	//	"github.com/garyburd/redigo/redis"
	//"github.com/Unknwon/goconfig"
	//	"strconv"
	//"bytes"
	//	"os"
	//	"bufio"
	//	"io"
	//	"strings"
	//m "model"
	//"utils"
	"strconv"
	"time"
)

//global
var LISTMAX int
var PAGEMAX int
var NAVMAX int

var CAT_INT_STR map[int]string

var CAT_STR_INT map[string]int

var CAT_INT_STRCN map[int]string


func CheckErr(err error) {
	if err != nil {
		log.Error("Error...", err)
	}
}

func IntToDateStr(d int64) string {
	tm := time.Unix(d, 0)
	ts := tm.Format("2006-01-02 15:04:05")
	return ts
}

func IntToStr(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}

func InitCateMap() {
	CAT_INT_STR = map[int]string{}
	CAT_INT_STRCN = map[int]string{}
	CAT_STR_INT = map[string]int{}

	CAT_INT_STR[0] = "all"
	CAT_INT_STR[1] = "video"
	CAT_INT_STR[2] = "torrent"
	CAT_INT_STR[3] = "soft"
	CAT_INT_STR[4] = "doc"
	CAT_INT_STR[5] = "music"
	CAT_INT_STR[6] = "picture"
	CAT_INT_STR[7] = "other"

	CAT_INT_STRCN[0] = "全部"
	CAT_INT_STRCN[1] = "视频"
	CAT_INT_STRCN[2] = "种子"
	CAT_INT_STRCN[3] = "软件"
	CAT_INT_STRCN[4] = "文档"
	CAT_INT_STRCN[5] = "音乐"
	CAT_INT_STRCN[6] = "图片"
	CAT_INT_STRCN[7] = "其他"

	CAT_STR_INT["all"] = 0
	CAT_STR_INT["video"] = 1
	CAT_STR_INT["torrent"] = 2
	CAT_STR_INT["soft"] = 3
	CAT_STR_INT["doc"] = 4
	CAT_STR_INT["music"] = 5
	CAT_STR_INT["picture"] = 6
	CAT_STR_INT["other"] = 7
}
