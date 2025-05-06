package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/scroll"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/joho/godotenv"
	"icomm/esintegration/models"
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
	}
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	searchResult, err := esClient.Search().Index(os.Getenv(("ES_INDEX"))).Request(&search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}).Scroll("30m").Do(context.TODO())

	if err != nil {
		panic(err)
	}

	hitsMetadata := searchResult.Hits
	scrollId := searchResult.ScrollId_

	for {
		if hitsMetadata.Total.Value == 0 {
			break
		}

		processData(hitsMetadata.Hits)

		scrollResponse, err := esClient.Scroll().Request(&scroll.Request{
			ScrollId: *scrollId,
			Scroll:   "1d",
		}).Do(context.TODO())
		if err != nil {
			panic(err)
		}
		hitsMetadata = scrollResponse.Hits
		scrollId = searchResult.ScrollId_
	}
}

func processData(data []types.Hit) {
	for _, hit := range data {
		var esDocument models.ES_Document
		if err := json.Unmarshal(hit.Source_, &esDocument); err != nil {
			panic(err)
		}
		fmt.Printf("Got document: %+v\n", esDocument)
	}
}
