package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	api "github.com/ksukhorukov/atlant/api"
)

const (
	address     = "localhost:31337"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	errorCheck(err)

	defer conn.Close()

	c := api.NewApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fetch_request, fetch_err := c.Fetch(ctx, &api.FetchRequest{Url: "http://localhost:3000/sequence.csv"})

	errorCheck(fetch_err)

	log.Printf("Imported: %d", fetch_request.GetCount())

	list_request, list_err := c.List(ctx, &api.ListRequest{
		Column: "price",
	 	Order: 1, //ascending, -1 means descending
	 	PageNumber: 1,
	 	ResultsPerPage: 60,
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