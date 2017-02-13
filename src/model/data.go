package model

import (
	//	"fmt"
	es "gopkg.in/olivere/elastic.v3"
	"encoding/json"
	"github.com/siddontang/go/log"
//	u "utils"
	"math/rand"
	"time"
)


var TotalShares int64
var TotalUsers int64
var TotalKeywords int64


func SearchShare(esclient *es.Client, query es.Query, page int, pagesize int, sort string)([]Share, int64) {
	searchResult := Search(esclient, "sharedata", query, page, pagesize, sort)
	if searchResult == nil {
		return nil, 0
	}

	shares := []Share{}
	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			sd := ShareData{}

			err := json.Unmarshal(*hit.Source, &sd)
			if err != nil {
				log.Error("Failed to read search result", err)
			}

			if hit.Highlight != nil && len(hit.Highlight) != 0 {
				if hl, found := hit.Highlight["title"]; found {
					sd.HTitle = hl[0]
				} else {
					sd.HTitle = sd.Title
				}
			} else {
				sd.HTitle = sd.Title
			}
			s := ShareDataToShare(sd)
			shares = append(shares, s)
		}
	} else {
		return nil, 0
	}

	return shares, searchResult.Hits.TotalHits
}

func SearchUser(esclient *es.Client, query es.Query, page int, pagesize int)([]User, int64) {
	searchResult := Search(esclient, "uinfo", query , page, pagesize, "")
	if searchResult == nil {
		return nil, 0
	}
	users := []User{}
	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			u := UserInfo{}

			err := json.Unmarshal(*hit.Source, &u)
			if err != nil {
				log.Error("Failed to read search result", err)
			}
			user := UserInfoToUser(u)
			users = append(users, user)
		}
	} else {
		return nil, 0
	}
	return users, searchResult.Hits.TotalHits
}


func SearchKeyword(esclient *es.Client, query es.Query, page int, pagesize int)([]Keyword, int64) {
	searchResult := Search(esclient, "keyword", query , page, pagesize, "count")
	if searchResult == nil {
		return nil, 0
	}
	keywords := []Keyword{}
	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			k := Keyword{}

			err := json.Unmarshal(*hit.Source, &k)
			if err != nil {
				log.Error("Failed to read search result", err)
			}

			keywords = append(keywords, k)

		}
	} else {
		return nil, 0
	}
	return keywords, searchResult.Hits.TotalHits
}



func Search(esclient *es.Client, index string,  query es.Query, page int, pagesize int, sort string) *es.SearchResult {
	// Specify highlighter
	hl := es.NewHighlight()
	hl = hl.Fields(es.NewHighlighterField("title"))
	hl = hl.PreTags("<mark>").PostTags("</mark>")
	hl = hl.Encoder("utf-8")

	start := (page - 1) * pagesize
	searchService := esclient.Search().
		Index(index).
		Highlight(hl).
		Query(query).
		From(start).Size(pagesize).
		Pretty(true)

	if sort != "" {
		searchService = searchService.Sort(sort, false)
	}

	searchResult, err := searchService.Do()                // execute
	if err != nil {
		log.Info(err)
		return nil
	}

	log.Info("Query took ", searchResult.TookInMillis, " msec")
	// Here's how you iterate through the search results with full control over each step.
	log.Info("Found a total of ", searchResult.Hits.TotalHits)
	return searchResult
}

func GetTotalShares(esclient *es.Client) int64 {
	query:= es.NewMatchAllQuery()
	var size int64
	_, size = SearchShare(esclient, query, 1, 1, "")
	return size
}

func GetTotalKeywords(esclient *es.Client) int64 {


	query:= es.NewMatchAllQuery()
	var size int64
	_, size = SearchKeyword(esclient, query, 1, 1)
	return size
}


func GetTotalUsers(esclient *es.Client) int64 {
	var size int64
	query:= es.NewMatchAllQuery()
	_, size = SearchUser(esclient, query, 1, 1)
	return size
}

func GenerateRandomShares(esclient *es.Client, category int, size int, keyword string) []Share{
	var start int
	if TotalShares < int64(size) {
		start = 1
	} else {
		max := int(TotalShares) / size
		rand.Seed(time.Now().UnixNano())
		if max <= 1 {
			max = 1
		}
		start = rand.Intn(max -1) + 1
	}


	boolQuery := es.NewBoolQuery()
	if keyword != "" {
		boolQuery.Should(es.NewQueryStringQuery(keyword))
		start = 2
	}
	query := es.NewTermQuery("search", 1)
	boolQuery.Should(query)

	if category != 0 {
		boolQuery.Must(es.NewTermQuery("category", category))
	}

	randomShares, _ := SearchShare(esclient, boolQuery, start, size, "")
	return randomShares
}


func GenerateRandomUsers(esclient *es.Client, size int) []User{
	var start int
	if TotalUsers < int64(size) {
		start = 1
	} else {
		max := int(TotalUsers) / size
		if max <= 1 {
			max = 1
		}
		rand.Seed(time.Now().UnixNano())
		start = rand.Intn(max)
	}

	query:= es.NewMatchAllQuery()
	randomUsers, _ := SearchUser(esclient, query, start, size)
	return randomUsers
}


func GenerateRandomKeywords(esclient *es.Client, size int) []Keyword{
	var start int
	if TotalKeywords < int64(size) {
		start = 1
	} else {
		max := int(TotalKeywords) / size
		rand.Seed(time.Now().UnixNano())
		if max <= 1 {
			max = 1
		}
		start = rand.Intn(max -1) + 1
	}

	query:= es.NewMatchAllQuery()
	randomKeywords, _ := SearchKeyword(esclient, query, start, size)
	return randomKeywords
}


func GenerateUserShares(esclient *es.Client, uk string, size int) []Share{
	query := es.NewTermQuery("uk", uk)
	userShares, _ := SearchShare(esclient, query, 1, size, "")
	return userShares
}
