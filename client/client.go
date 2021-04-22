package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	api "github.com/ksukhorukov/atlant/api"
)

var server_address string
var fetch_url string
var show_help bool

const DEFAULT_SERVER_ADDRESS = "localhost:55555"
const DEFAULT_FETCH_URL = "http://localhost:3000/products.csv"

func main() {
	systemParams()

	if show_help {
		usage()
		os.Exit(1)
	}

	conn, err := grpc.Dial(server_address, grpc.WithInsecure(), grpc.WithBlock())

	errorCheck(err)

	defer conn.Close()

	c := api.NewApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fetch_request, fetch_err := c.Fetch(ctx, &api.FetchRequest{Url: fetch_url})

	errorCheck(fetch_err)

	log.Printf("Imported: %d", fetch_request.GetCount())

	list_request, list_err := c.List(ctx, &api.ListRequest{
		Column:         "price",
		Order:          1, //ascending, -1 means descending
		PageNumber:     1,
		ResultsPerPage: 50,
	})

	errorCheck(list_err)

	results := list_request.GetResults()

	for i := 0; i < len(results); i++ {
		record := results[i]

		log.Printf("Product: %s, Price: %f, Times price changed: %d, Request time: %v\n",
			record.GetProduct(),
			record.GetPrice(),
			record.GetTimespricechanged(),
			time.Unix(record.GetRequesttime(), 0))
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func systemParams() {
	flag.StringVar(&server_address, "server", DEFAULT_SERVER_ADDRESS, "Address of our server")
	flag.StringVar(&fetch_url, "url", DEFAULT_FETCH_URL, "CSV file URL")
	flag.BoolVar(&show_help, "help", false, "Help center")
	flag.Parse()
}

func usage() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("%s --server=localhost:5555 --url=http://localhost:3000/products.csv\n\n", os.Args[0])
}
