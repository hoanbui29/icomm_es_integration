package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/scroll"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/joho/godotenv"
	"icomm/esintegration/models"
	"net/http"
	"os"
	"strings"
)

type esData struct {
	ID string `json:"id"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	esAddresses := strings.Split(os.Getenv("ES_ADDRESSES"), ",")
	fmt.Printf("ES_ADDRESSES: %s\n", esAddresses)

	cfg := elasticsearch.Config{
		Addresses: esAddresses,
		Username:  os.Getenv("ES_USERNAME"),
		Password:  os.Getenv("ES_PASSWORD"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	searchResult, err := esClient.Search().Index(os.Getenv(("ES_INDEX"))).Request(&search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}).Size(1000).Scroll("30m").Do(context.TODO())

	if err != nil {
		panic(err)
	}

	hitsMetadata := searchResult.Hits
	scrollId := searchResult.ScrollId_
	count := 0

	for {
		if len(hitsMetadata.Hits) == 0 {
			break
		}

		processData(hitsMetadata.Hits)
		count += len(hitsMetadata.Hits)
		fmt.Printf("Processed %d documents\n", count)

		scrollResponse, err := esClient.Scroll().Request(&scroll.Request{
			ScrollId: *scrollId,
			Scroll:   "1d",
		}).Do(context.TODO())
		if err != nil {
			panic(err)
		}
		hitsMetadata = scrollResponse.Hits
		scrollId = scrollResponse.ScrollId_
	}
}

func processData(data []types.Hit) {
	for _, hit := range data {
		var esDocument models.ES_Document
		if err := json.Unmarshal(hit.Source_, &esDocument); err != nil {
			panic(err)
		}
	}
}
