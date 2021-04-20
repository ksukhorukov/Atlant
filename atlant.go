package main

import (
	 api "github.com/ksukhorukov/atlant/proto"

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
	"os"
)

const (
	port = ":31337"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	api.UnimplementedApiServer
}

const ERROR_INCORRECT_STRUCTURE = "Incorrect CSV file structure"
const ERROR_INCORRECT_HEADERS 	= "Incorrect CSV file headers"
const ERROR_INCORRECT_FILE_TYPE = "Incorrect file type"
const DOWNLOAD_DIRECTORY = "./tmp"

const DB_NAME = "atlant"
const DB_COLLECTION_NAME = "products"

type Record struct {
	Product string 
	Price float64
	TimesPriceChanged int
	RequestTime time.Time
}


func main() {
	mng_context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, collection := initMongo(mng_context)

	defer client.Disconnect(mng_context)



	// args := os.Args

	// file_url := args[1]

	// file_path := randomFile(DOWNLOAD_DIRECTORY, 64)

	// err := downloadFile(file_url, file_path)

	// errorCheck(err)

	// fmt.Printf("[+] Starting to parse %s\n", file_path)

	// timestamp := time.Now()

	// parseCSV(file_path, collection, mng_context, timestamp)

	// err = deleteFile(file_path)

	// errorCheck(err)

	// column := "price" 
	// //filter := "ascending"
	// filter := "descending"

	// var results []Record

	// results = search(1, 10, column, filter, collection, mng_context)

	// printResults(results)
}

func search(page int, per_page int, column string, filter string, collection mongo.Collection, mng_context context.Context) []Record {
	var results []Record
	var order int

	if(filter == "ascending") {
		order = 1
	} else {
		order = -1
	}

	opts := options.Find().SetSort(bson.D{{column, order}})

	cursor, err := collection.Find(mng_context, bson.D{{}}, opts)	

	errorCheck(err)

	err = cursor.All(mng_context, &results)

	errorCheck(err)

	cursor_index := getCursorIndex(page, per_page, len(results))

	return results[cursor_index:per_page]
}

func getCursorIndex(page int, per_page int, length int) int {
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

func parseCSV(file_path string, collection mongo.Collection, mng_context context.Context, timestamp time.Time) {
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

		saveResults(collection, mng_context, product, price, timestamp)
	}
}

func saveResults(collection mongo.Collection, mng_context context.Context, product string, price float64, timestamp time.Time) {
		var result Record
		
		err := collection.FindOne(mng_context, bson.D{{"product", product}}).Decode(&result)

		if err != nil { // nothing found
			record := Record{product, price, 0, timestamp}
			_, err = collection.InsertOne(mng_context, record)

			errorCheck(err)
		} else { // need to update existing record
			if(result.Price == price) { // exit if nothing changed
 				return 
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
  	}

  	fmt.Printf("[+] %s successfully updated!\n", product)
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
