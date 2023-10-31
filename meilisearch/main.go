package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/meilisearch/meilisearch-go"
)

// object yang akan kita tambahkan ke meilisearch
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
}

// func create(index *meilisearch.Index, products []Product) (err error) {
// 	// change request into slice of map
// 	var req = []map[string]interface{}{}

// 	reqByte, err := json.Marshal(products)
// 	if err != nil {
// 		return
// 	}

// 	err = json.Unmarshal(reqByte, &req)
// 	if err != nil {
// 		return
// 	}

// 	// "id" adalah primary key yang akan kita gunakan
// 	task, err := index.AddDocuments(req, "id")
// 	if err != nil {
// 		return
// 	}

// 	// id task untuk melihat status process
// 	log.Println("task id:", task.TaskUID)
// 	return
// }

type SearchProductModel struct {
	Query  string   `json:"query"`
	Facets []string `json:"facets"`
	Filter string   `json:"filter"`
	Sort   []string `json:"sort"`

	Pagination
}

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func get(index *meilisearch.Index, req SearchProductModel) (products []Product, err error) {
	// search data
	response, err := index.Search(req.Query, &meilisearch.SearchRequest{
		HitsPerPage: int64(req.Limit),
		Page:        int64(req.Page),
		Filter:      req.Filter,
		Facets:      req.Facets,
	})
	if err != nil {
		return
	}

	// response.hits adalah kumpulan data dari hasil proses
	// pencarian
	hitByte, err := json.Marshal(response.Hits)
	if err != nil {
		return
	}

	err = json.Unmarshal(hitByte, &products)
	if err != nil {
		return
	}

	// response.FacetDistribution adalah total
	// data dari hasil pengelompokan
	log.Printf("%+v\n", response.FacetDistribution)
	return
}

func main() {
	client, err := connectMeilisearch()
	if err != nil {
		panic(err)
	}

	// set index products
	index := client.Index("products")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	log.Println(index)

	// // set data
	// product := Product{
	// 	Name:        "Mouse Logitech",
	// 	Description: "Mouse logitech terbaru dengan beragam pilihan warna",
	// 	Category:    "Mouse",
	// 	Price:       200_000,
	// 	Id:          1,
	// }

	// // tambahkan ke dalam array.
	// // request body yang dibutuhkan merupakan sebuah array
	// products := []Product{
	// 	product,
	// }

	// err = create(index, products)
	// if err != nil {
	// 	panic(err)
	// }

	myProducts, err := get(index, SearchProductModel{
		Query:  "logitech",           // kita buat typo sedikit
		Facets: []string{"category"}, // lalu tampilkan total data berdasarkan category
		Pagination: Pagination{
			Limit: 100,
			Page:  1,
		},
	})
	log.Println(myProducts)
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range myProducts {
		// hanya untuk pembatas
		div := strings.Repeat("=", 11)
		log.Println(div, "[PRODUCT]", div)
		log.Printf("%+v\n", p)
		log.Println(div + div + div)
	}
}

func connectMeilisearch() (client *meilisearch.Client, err error) {
	// init meilisearch client
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",                       // setup meilisearch host
		APIKey: "IZW1vyBmNtQ5TlcXeSPBXlqcuKQCsXgPt9TzaAe1T_I", // setup meilisearch api key
	})

	// validate is client null or not
	if client == nil {
		return nil, errors.New("error when try to connect to meilisearch")
	}

	return
}
