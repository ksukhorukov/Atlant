package main

import (
	 api "github.com/ksukhorukov/atlant/api"

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
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)


const (
	DEFAULT_SERVER_ADDRESS = "127.0.0.1"
	DEFAULT_SERVER_PORT = 55555
	
	DEFAULT_MONGO_ADDRESS = "mongo" //"127.0.0.1"
	DEFAULT_MONGO_PORT = 27017

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
	TimesPriceChanged int64
	RequestTime int64
}

type saver func(mongo.Collection, context.Context, string, float64, int64) bool

var server_address = DEFAULT_SERVER_ADDRESS
var server_port = DEFAULT_SERVER_PORT

var mongo_address = DEFAULT_MONGO_ADDRESS
var mongo_port = DEFAULT_MONGO_PORT

var show_help = false


func (s *server) Fetch(ctx context.Context, in *api.FetchRequest) (*api.FetchResponse, error) {
	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	log.Printf("Received: %v", in.GetUrl())

	file_path := RandomFile(DOWNLOAD_DIRECTORY, 64)

	err := DownloadFile(in.GetUrl(), file_path)

	ErrorCheck(err)

	fmt.Printf("[+] Starting to parse %s\n", file_path)

	timestamp := time.Now().Unix()

	saver := SaveResults

	var count int64

	count, err = ParseCSV(file_path, saver, collection, mng_context, timestamp)

	ErrorCheck(err)

	err = DeleteFile(file_path)

	ErrorCheck(err)

	return &api.FetchResponse{Count: count}, nil
}

func (s *server) List(ctx context.Context, in *api.ListRequest) (*api.ListResponse, error) {
	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	column := in.GetColumn()
	order := in.GetOrder()
	page := in.GetPageNumber()
	results_per_page := in.GetResultsPerPage()

	log.Printf("Received. Column: %v, Order: %v, PageNumber: %v, ResultsPerPage: %v", 
		column, order, page, results_per_page)


	var results []api.Result

	results = Search(int64(page), int64(results_per_page), column, int32(order), collection, mng_context)

	data_size := len(results)
	data := make([]*api.Result, data_size)

	for i := 0; i < data_size; i++ {
		data[i] = &results[i]
	}

	return &api.ListResponse{Results: data}, nil
}


func main() {
	SystemParams()

	if show_help {
		Usage()
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", SocketAddress())
	
	ErrorCheck(err)

	defer lis.Close()

	s := grpc.NewServer()
	
	api.RegisterApiServer(s, &server{})
	
	err = s.Serve(lis)

	ErrorCheck(err)
}

func Search(page int64, per_page int64, column string, order int32, collection mongo.Collection, mng_context context.Context) []api.Result {
	var results []api.Result
	
	opts := options.Find().SetSort(bson.D{{column, order}})

	cursor, err := collection.Find(mng_context, bson.D{{}}, opts)	

	ErrorCheck(err)

	err = cursor.All(mng_context, &results)

	ErrorCheck(err)

	results_size := int64(len(results))

	start, end := GetCursorRange(page, per_page, results_size)

	return results[start:end]
}

func GetCursorRange(page int64, per_page int64, length int64) (int64, int64) {
	var start int64

	if length < per_page {
		return 0, length
	}

	if page == 1 || page == 0 {
		start = 0
	}

	if length <= 0 {
		return 0, 0
	}

	if page > 1 {
		page = page - 1

		if per_page * page >= length {
			start = 0
		} else {
			start = page * per_page 
		}
	} else if page < 0 {
		if (page * -1) * per_page < length {
			start = length + (page * per_page)	
		}
	}

	if start + per_page > length {
		return start, length
	}

	return start, start + per_page
}


func PrintResults(results []Record) {
	for _, result := range results {
		fmt.Printf("%s %f %s %d\n", result.Product, result.Price, time.Unix(result.RequestTime, 0), result.TimesPriceChanged)
	}	
}

func InitMongo(mng_context context.Context)(mongo.Client, mongo.Collection)  {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoAddress()))
	
	ErrorCheck(err)

	err = client.Connect(mng_context)

	ErrorCheck(err)

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)

	err = client.Ping(mng_context, nil)

	ErrorCheck(err)

	fmt.Printf("[+] Connected to MongoDB\n")

	return *client, *collection
}

func ParseCSV(file_path string, saver saver, collection mongo.Collection, mng_context context.Context, timestamp int64)(int64, error) {
	var counter int64

	counter = 0

	file, err := os.Open(file_path)
	
	if err != nil {
		return 0, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	headers, err := reader.Read()

	if err != nil {
		return 0, err
	}

	err = CheckHeaders(headers)

	if err != nil {
		return 0, err
	}

	records, err := reader.ReadAll()
	
	if err != nil {
		return 0, err
	}

	for _, record := range records {
		err := CheckStructure(record)

		if err != nil {
			return counter, err
		}

		product := record[0]
		price, err := ConvertStringToFloat(record[1])

		if err != nil {
			return counter, err
		}

		if(saver(collection, mng_context, product, price, timestamp)) {
			counter += 1
		}
	}

	return counter, nil
}

func SaveResults(collection mongo.Collection, mng_context context.Context, product string, price float64, timestamp int64)bool {
		var result Record

		saved := false
		
		err := collection.FindOne(mng_context, bson.D{{"product", product}}).Decode(&result)

		if err != nil { // nothing found
			record := Record{product, price, 0, timestamp}
			_, err = collection.InsertOne(mng_context, record)

			ErrorCheck(err)

			saved = true
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

  	  ErrorCheck(err)

  	  saved = true
  	}

  	return saved
}

func ConvertStringToFloat(str string) (float64, error) {
	  fnumber, err := strconv.ParseFloat(str, 64)

	  return fnumber, err
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckHeaders(headers []string) error {
	if headers[0] != "PRODUCT NAME" || headers[1] != "PRICE" {
		return fmt.Errorf("%s\n", ERROR_INCORRECT_HEADERS)
	}

	return nil
}

func CheckStructure(record []string) error {
	if len(record) != 2 {
		return fmt.Errorf("%s\n", ERROR_INCORRECT_STRUCTURE)
	}

	return nil
}

func DownloadFile(url string, filepath string) error {
	resp, err := http.Get(url)

	//ErrorCheck(err)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	mime := mimetype.Detect(body)

	err = CheckMimeType(mime)

	//ErrorCheck(err)

	if err != nil {
		return err
	}

	_ = os.Mkdir(DOWNLOAD_DIRECTORY, 0777)

	// ErrorCheck(err)

	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()

	// ErrorCheck(err)

	if err != nil {
		return err
	}

	_, err = out.Write(body)

	return err
}

func RandomFile(dir string, n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("%s/%s.txt", dir, string(s))
}

func DeleteFile(file_path string) error {
	err := os.Remove(file_path)

	return err
}

func CheckMimeType(mime *mimetype.MIME) error {
	if mime.Is("text/plain") == false {
		return fmt.Errorf("%s", ERROR_INCORRECT_FILE_TYPE)
	}

	return nil
}

func SystemParams() {
	flag.StringVar(&server_address, "host", DEFAULT_SERVER_ADDRESS, "Address of our server")
	flag.IntVar(&server_port, "port", DEFAULT_SERVER_PORT, "Service port number")

	flag.StringVar(&mongo_address, "mongo_address", DEFAULT_MONGO_ADDRESS, "Address of MongoDB server")
	flag.IntVar(&mongo_port, "mongo_port", DEFAULT_MONGO_PORT, "MongoDB port number")

	flag.BoolVar(&show_help, "help", false, "Help center")

	flag.Parse()
}

func Usage() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("%s --host=0.0.0.0 --port=55555 --mongo_address=192.168.0.100 --mongo_port=27017\n\n", os.Args[0])

	fmt.Printf("Default settings:\n\n")
	fmt.Printf("Host: %s\n", DEFAULT_SERVER_ADDRESS)
	fmt.Printf("Port: %d\n", DEFAULT_SERVER_PORT)
	fmt.Printf("MongoDB address: %s\n", DEFAULT_MONGO_ADDRESS)
	fmt.Printf("MongoDB port: %d\n", DEFAULT_MONGO_PORT)
}

func SocketAddress() string {
	return fmt.Sprintf("%s:%s", server_address, strconv.Itoa(server_port))
}

func MongoAddress() string {
	return fmt.Sprintf("mongodb://%s:%s", mongo_address, strconv.Itoa(mongo_port))
}