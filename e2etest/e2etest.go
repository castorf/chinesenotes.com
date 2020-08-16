/**
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

/**
 * End-to-end test program
 */

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alexamies/chinesenotes-go/applog"
	"github.com/alexamies/chinesenotes-go/dictionary"
	"github.com/alexamies/chinesenotes-go/dicttypes"
	"github.com/alexamies/chinesenotes-go/find"
	"github.com/alexamies/chinesenotes-go/fulltext"
	"github.com/alexamies/chinesenotes-go/webconfig"
	"log"
	"net/http"
	"os"
)

const STATIC_DIR string = "./static"

var (
  database *sql.DB
	dictSearcher *dictionary.Searcher
	parser find.QueryParser
	wdict map[string]dicttypes.Word
)

func init() {
	appInit()
}

func appInit() {
	applog.Info("e2etest.appInit Initializing e2etest")
	var err error
	database, err = initDBCon()
	if err != nil {
		applog.Errorf("e2etest.appInit unable to initialize databsae connection: %v", err)
	}
	ctx := context.Background()
	dictSearcher = dictionary.NewSearcher(ctx, database)
	applog.Errorf("e2etest.appInit dictSearcher.DatabaseInitialized(): %v",
			dictSearcher.DatabaseInitialized())
	wdict, err = dictionary.LoadDict(ctx, database)
	if err != nil {
		applog.Errorf("e2etest.appInit unable to load dictionary: %v", err)
	}
	parser = find.MakeQueryParser(wdict)
}

func initDBCon() (*sql.DB, error) {
	conString := webconfig.DBConfig()
	dbPool, err := sql.Open("mysql", conString)
	if err != nil {
		return nil, fmt.Errorf("sql.Open with with conn string %s: %v", conString, err)
	}
	return dbPool, nil
}

// Finds documents matching the given query with search in text body
func findAdvanced(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findAdvanced, enter")
	findDocs(response, request, true)
}

// Finds documents matching the given query
func findDocs(response http.ResponseWriter, request *http.Request, advanced bool) {
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	applog.Infof("e2etest.findDocs, query: %s", query)
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	} else {
		query := queryString["text"]
		if len(query) > 0 {
			q = query[0]
		}
	}

	var results *find.QueryResults
	var err error

	c := queryString["collection"]
	ctx := context.Background()
	if (len(c) > 0) && (c[0] == "xiyouji.html")  {
		applog.Error("main.findDocs mock data for xiyouji.html")
		col0 := find.Collection{
			GlossFile: "xiyouji.html",
			Title: "Journey to the West 《西遊記》",
		}
		col := []find.Collection{col0}
		ft := fulltext.MatchingText{
			Snippet: "詩曰：",
			LongestMatch: "詩",
			ExactMatch: true,
		}
		doc0 := find.Document{
			GlossFile: "xiyouji/xiyouji001.html",
			Title: "第一回 Chapter 1",
			CollectionFile: "xiyouji.html",
			CollectionTitle: "Journey to the West 《西遊記》",
			ContainsWords: "詩",
			ContainsBigrams: "",
			SimTitle: 0.0,
			SimWords: 1.0,
			SimBigram: 0.0,
			SimBitVector: 1.0,
			Similarity: 1.0,
			ContainsTerms: []string{"詩"},
			MatchDetails: ft,
		}
		doc := []find.Document{doc0}
		senses0 := []dicttypes.WordSense{}
		sense1 := dicttypes.WordSense{
			Id: 5925,
			HeadwordId: 5925,
			Simplified: "诗",
			Traditional: "詩",
			Pinyin: "shī",
			English: "poem",
			Notes: "",
		}
		senses1 := []dicttypes.WordSense{sense1}
		entry1 := dicttypes.Word{
			Simplified: "诗",
			Traditional: "詩",
			Pinyin: "shī",
			HeadwordId: 5925,
			Senses: senses1,
		}
		ts1 := find.TextSegment{
			QueryText: "詩",
			DictEntry: entry1,
			Senses: senses0,
		}
		terms1 := []find.TextSegment{ts1}
		results = &find.QueryResults{q, "xiyouji.html", 1, 1, col, doc, terms1}
	} else if (len(c) > 0) && (c[0] != "") {
		applog.Infof("e2etest.findDocs finding docs for collection %s", c[0])
		results, err = find.FindDocumentsInCol(ctx, dictSearcher, parser, q, c[0])
	} else {
		applog.Infof("e2etest.findDocs find with advanced = %v ", advanced)
		results, err = find.FindDocuments(ctx, dictSearcher, parser, q, advanced)
	}

	if err != nil {
		applog.Errorf("e2etest.findDocs Error searching docs, %v", err)
		http.Error(response, "Error searching docs",
			http.StatusInternalServerError)
		return
	}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("e2etest.findDocs error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results",
			http.StatusInternalServerError)
	} else {
		if (q != "hello" && q != "Eight" ) { // Health check monitoring probe
			applog.Infof("e2etest.findDocs, results: %s", string(resultsJson))
		}
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

// Finds documents matching the given query
func findHandler(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findHandler, enter:")
	findDocs(response, request, false)
}

// Finds terms matching the given query with a substring match
func findSubstring(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findSubstring, enter")
	sense := dicttypes.WordSense{
			Id: 62084,
			HeadwordId: 62084,
			Simplified: "同床异梦",
			Traditional: "同床異夢",
			Pinyin: "tóng chuáng yì mèng",
			English: "to share the same bed with different dreams",
			Notes: "",
	}
	senses := []dicttypes.WordSense{sense}
	word := dicttypes.Word{
		Simplified: "同床异梦",
		Traditional: "同床異夢",
		Pinyin: "tóng chuáng yì mèng",
		HeadwordId: 62084,
		Senses: senses,
	}
	words := []dicttypes.Word{word}
	results := dictionary.Results{words}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.findSubstring error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results",
			http.StatusInternalServerError)
	} else {
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

func main() {
	log.Print("End-to-end test server started")
	http.HandleFunc("/find/", findHandler)
	http.HandleFunc("/findadvanced/", findAdvanced)
	http.HandleFunc("/findsubstring", findSubstring)
	http.Handle("/", http.FileServer(http.Dir(STATIC_DIR)))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}