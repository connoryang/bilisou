package model

import (
//	"fmt"
	es "gopkg.in/olivere/elastic.v3"
//	"encoding/json"
//	"github.com/siddontang/go/log"
//	u "utils"
)


func GenerateSearchPageVar(esclient *es.Client, category int, keyword string, page int) *PageVar {
	if page <= 0 {
		return nil
	}

	pv := PageVar{}
	pv.Type = "search"
	pv.Keyword = keyword

	boolQuery := es.NewBoolQuery()
	query := es.NewQueryStringQuery(keyword)
	boolQuery.Must(query)
	if category != 0 {
		boolQuery.Must(es.NewTermQuery("category", category))
	}
	var size int64
	pv.SearchShares, size = SearchShare(esclient, boolQuery, page, 20, "")
	//log.Info(pv.SearchShares

	if len(pv.SearchShares) == 0 {
		//return nil
		pv.Type = "lost"
	}

	//pv.End = int(size)
	pv.End = int(size) / 20 + 1
	pv.Current = page

	SetBA(&pv)
	SetCategory(&pv, category)

	pv.RandomUsers = GenerateRandomUsers(esclient, 24)
	pv.Keywords = GenerateRandomKeywords(esclient, 30)

	return &pv

}
