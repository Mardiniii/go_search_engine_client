package main

// Elastic search client
import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/olivere/elastic"
)

const (
	indexName    = "pages"
	indexMapping = `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"page":{
				"properties":{
					"title": {
						"type":"text"
					},
					"description": {
						"type":"text"
					},
					"body": {
						"type":"text"
					},
					"url": {
						"type":"text"
					}
				}
			}
		}
	}`
)

var client *elastic.Client

// NewElasticSearchClient returns an elastic seach client
func NewElasticSearchClient() *elastic.Client {
	var err error

	// Create a new elastic client
	client, err = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return client
}

// ExistsIndex checks if the given index exists or not
func ExistsIndex(i string) bool {
	// Check if index exists
	exists, err := client.IndexExists(i).Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

// CreateIndex creates a new index
func CreateIndex(i string) {
	createIndex, err := client.CreateIndex(indexName).
		Body(indexMapping).
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	if !createIndex.Acknowledged {
		log.Println("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
}

// SearchContent returns the results for a given query
func SearchContent(input string) []Page {
	pages := []Page{}

	ctx := context.Background()
	// Search for a page in the database using multi match query
	q := elastic.NewMultiMatchQuery(input, "title", "description", "body", "url").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := client.Search().
		Index(indexName).
		Pretty(true).
		Sort("_score", false).
		Query(q).
		Do(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var ttyp Page
	for _, page := range result.Each(reflect.TypeOf(ttyp)) {
		p := page.(Page)
		pages = append(pages, p)
	}

	return pages
}
