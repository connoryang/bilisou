package model

import (
		"fmt"
	//	u "utils"
	"github.com/siddontang/go/log"
	"database/sql"
//	es "gopkg.in/olivere/elastic.v3"

)



func KeywordHit(db *sql.DB, keyword string) {
	s := "select count from keyword where keyword = '%s'"
	s = fmt.Sprintf(s, keyword)
	log.Info(s)
	rows, _ := db.Query(s)
	if rows.Next() {
		var count int64
		rows.Scan(&count)
		count = count + 1;
		us := "update keyword set count = %d  where keyword = '%s'"
		us = fmt.Sprintf(us, count, keyword)
		log.Info(us)
		stmt, _ := db.Prepare(us)
		stmt.Exec()
		stmt.Close()
	} else {
		us := "insert into keyword(keyword) values('%s')"
		us = fmt.Sprintf(us, keyword)
		log.Info(us)
		stmt, _ := db.Prepare(us)
		stmt.Exec()
		stmt.Close()
	}
}
