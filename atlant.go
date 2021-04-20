package main

import (
	 api "github.com/ksukhorukov/atlant/proto"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"encoding/csv"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	port = ":31337"
	
	ERROR_INCORRECT_STRUCTURE = "Incorrect CSV file structure"
	ERROR_INCORRECT_HEADERS 	= "Incorrect CSV file headers"
	ERROR_INCORRECT_FILE_TYPE = "Incorrect file type"
	DOWNLOAD_DIRECTORY = "./tmp"

	DB_NAME = "atlant"
	DB_COLLECTION_NAME = "products"
)

type server struct {
	api.UnimplementedApiServer
}

type Record struct {
	Product string 
	Price float64
	TimesPriceChanged int
	RequestTime time.Time
}

func (s *server) Fetch(ctx context.Context, in *api.FetchRequest) (*api.FetchResponse, error) {
	mng_context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, collection := initMongo(mng_context)

	defer client.Disconnect(mng_context)

	log.Printf("Received: %v", in.GetUrl())

	file_path := randomFile(DOWNLOAD_DIRECTORY, 64)

	err := downloadFile(in.GetUrl(), file_path)

	errorCheck(err)

	fmt.Printf("[+] Starting to parse %s\n", file_path)

	timestamp := time.Now()

	count := parseCSV(file_path, collection, mng_context, timestamp)

	err = deleteFile(file_path)

	errorCheck(err)

	return &api.FetchResponse{Count: count}, nil
}

func (s *server) List(ctx context.Context, in *api.ListRequest) (*api.ListResponse, error) {
	mng_context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, collection := initMongo(mng_context)

	defer client.Disconnect(mng_context)

	column := in.GetColumn()
	order := in.GetOrder()
	page := in.GetPageNumber()
	results_per_page := in.GetResultsPerPage()

	log.Printf("Received. Column: %v, Order: %v, PageNumber: %v, ResultsPerPage: %v", 
		column, order, page, results_per_page)


	var results []api.Result

	results = search(page, results_per_page, column, order, collection, mng_context)

	return &api.ListResponse{Results: results}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	
	errorCheck(err)

	s := grpc.NewServer()
	
	api.RegisterApiServer(s, &server{})
	
	err = s.Serve(lis)

	errorCheck(err)
}

func search(page int32, per_page int32, column string, order int32, collection mongo.Collection, mng_context context.Context) []api.Result {
	var results []api.Result
	
	opts := options.Find().SetSort(bson.D{{column, order}})

	cursor, err := collection.Find(mng_context, bson.D{{}}, opts)	

	errorCheck(err)

	err = cursor.All(mng_context, &results)

	errorCheck(err)

	cursor_index := getCursorIndex(page, per_page, int32(len(results)))

	return results[cursor_index:per_page]
}

func getCursorIndex(page int32, per_page int32, length int32) int32 {
	if(page == 1) {
		return 0
	}

	if(page > 1) {
		return (page * per_page) - 1
	}

	if(page < 0 && ((page * -1) * per_page <= length)) {
		return (length + (page * per_page)) - 1
	}

	return 0
}

func printResults(results []Record) {
	for _, result := range results {
		fmt.Printf("%s %f %s %d\n", result.Product, result.Price, result.RequestTime, result.TimesPriceChanged)
	}	
}

func initMongo(mng_context context.Context)(mongo.Client, mongo.Collection)  {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	
	errorCheck(err)

	err = client.Connect(mng_context)

	errorCheck(err)

	fmt.Printf("MONGO CLIENT TYPE: %T\n", client)

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)

	err = client.Ping(mng_context, nil)

	errorCheck(err)

	fmt.Printf("[+] Connected to MongoDB\n")

	return *client, *collection
}

func parseCSV(file_path string, collection mongo.Collection, mng_context context.Context, timestamp time.Time) int32 {
	var counter int32

	counter = 0

	file, err := os.Open(file_path)
	errorCheck(err)

	reader := csv.NewReader(file)
	reader.Comma = ';'

	headers, err := reader.Read()
	errorCheck(err)
	checkHeaders(headers)

	fmt.Printf("Headers: %s, %s\n", headers[0], headers[1])

	records, err := reader.ReadAll()
	errorCheck(err)

	for _, record := range records {
		checkStructure(record)

		product := record[0]
		price, err := convertStringToFloat(record[1])

		errorCheck(err)

		fmt.Printf("Product: %s, Price: %f\n", product, price)

		if(saveResults(collection, mng_context, product, price, timestamp)) {
			counter += 1
		}
	}

	return counter
}

func saveResults(collection mongo.Collection, mng_context context.Context, product string, price float64, timestamp time.Time) bool {
		var result Record

		saved := false
		
		err := collection.FindOne(mng_context, bson.D{{"product", product}}).Decode(&result)

		if err != nil { // nothing found
			record := Record{product, price, 0, timestamp}
			_, err = collection.InsertOne(mng_context, record)

			errorCheck(err)
		} else { // need to update existing record
			if(result.Price == price) { // exit if nothing changed
 				return false
 			}

			filter := bson.D{{"product", product}}
 
 			update := bson.D{
    		{"$set", bson.D{
    			{"price", price},
	        {"timespricechanged", result.TimesPriceChanged + 1},
	        {"requesttime", timestamp},
  	  	}},
  	  }

  	  _, err = collection.UpdateOne(mng_context, filter, update)

  	  errorCheck(err)

  	  saved = true
  	}

  	fmt.Printf("[+] %s successfully updated!\n", product)

  	return saved
}

func convertStringToFloat(str string) (float64, error) {
	  fnumber, err := strconv.ParseFloat(str, 64)

	  return fnumber, err
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkHeaders(headers []string) {
	if headers[0] != "PRODUCT NAME" || headers[1] != "PRICE" {
		log.Fatal(ERROR_INCORRECT_HEADERS)
	}
}

func checkStructure(record []string) {
	if len(record) != 2 {
		log.Fatal(ERROR_INCORRECT_STRUCTURE)
	}
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)

	errorCheck(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	mime := mimetype.Detect(body)

	err = checkMimeType(mime)

	errorCheck(err)

	_ = os.Mkdir(DOWNLOAD_DIRECTORY, 0777)

	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()

	errorCheck(err)

	_, err = out.Write(body)

	return err
}

func randomFile(dir string, n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("%s/%s.txt", dir, string(s))
}

func deleteFile(file_path string) error {
	err := os.Remove(file_path)

	return err
}

func checkMimeType(mime *mimetype.MIME) error {
	if mime.Is("text/plain") == false {
		return fmt.Errorf("%s", ERROR_INCORRECT_FILE_TYPE)
	}

	return nil
}
