package elasticbench

import (
	"../../internal/benchmarkutils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"strings"
)

var DocumentIds []string
var SearchTook string

func SearchAll() {
	//var b strings.Builder
	var mapResp map[string]interface{}
	//var QueryAll = `{"size": 10, "query": {"match_all": {}}}`
	response := SearchIndex()
	if err := json.NewDecoder(response.Body).Decode(&mapResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	SearchTook = fmt.Sprintf("%v", mapResp["took"])
	log.Printf("Search took time: %v", SearchTook)
	for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		DocumentIds = append(DocumentIds, fmt.Sprintf("%v", doc["_id"]))
	}
}

func SearchIndex() *esapi.Response {
	var b strings.Builder
	//var mapResp map[string]interface{}
	var QueryAll = `{"size": 10, "query": {"match_all": {}}}`
	b.WriteString(QueryAll)
	read := strings.NewReader(b.String())
	es := Connect()
	response, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(IndexName),
		es.Search.WithBody(read),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	benchmarkutils.ErrorCheck(err)
	return response
}
