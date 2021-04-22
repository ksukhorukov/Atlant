package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"testing"
	"regexp"
	"io/ioutil"
	"fmt"
	"os"
)

func TestMongoAddress(t *testing.T) {
	mongo_url := fmt.Sprintf("mongodb://%s:%d", mongo_address, mongo_port)
	
	func_result := MongoAddress()

	if(func_result != mongo_url) {
		t.Errorf("%s != %s\n", func_result, mongo_url)
	}
}

func TestSocketAddress(t *testing.T) {
	server_socket := fmt.Sprintf("%s:%d", server_address, server_port)

	func_result := SocketAddress()

	if(server_socket != func_result) {
		t.Errorf("%s != %s\n", func_result, server_socket)	
	}
}

func TestCheckMimeTypeWithCorrectMimeType(t *testing.T) {
	  data, err := ioutil.ReadFile("../samples/sample.csv")

    if err != nil {
        t.Errorf("Error reading sample file: %v\n", err)
        return
    }

    mime := mimetype.Detect(data)

    err = CheckMimeType(mime)

    if(err != nil) {
    	t.Errorf("Incorrect mimetype detection: %v", err)
    }
}

func TestCheckMimeTypeWithIncorrectMimeType(t *testing.T) {
	  data, err := ioutil.ReadFile("../samples/golang.png")

    if err != nil {
        t.Errorf("Error reading sample file: %v\n", err)
        return
    }

    mime := mimetype.Detect(data)

    err = CheckMimeType(mime)

    if(err == nil) {
    	t.Errorf("Incorrect mimetype detection: %v", err)
    }
}

func TestDeleteCorrectFile(t *testing.T) {
		tmp_file_path := "../samples/file_for_test_purposes.txt"
    tmpFile, err := os.Create(tmp_file_path)

    if err != nil {
        t.Errorf("Cannot create file for test purposes\n")
        return
    }

   	tmpFile.Close()

   	err = DeleteFile(tmp_file_path)

   	if err != nil {
   		t.Errorf("Cannot delete tmp file: %v", err)
   	}
}

func TestDeleteIncorrectFile(t *testing.T) {
		tmp_file_path := "../samples/file_for_test_purposes.txt"

   	err := DeleteFile(tmp_file_path)

   	if err == nil {
   		t.Errorf("Delete file function not working\n")
   	}
}

func TestRandomFileNameGeneration(t *testing.T) {
	random_file_name := RandomFile("../samples", 64)

	filename_pattern := regexp.MustCompile(`\.\.\/samples\/([0-9a-zA-Z]){64}\.txt$`)

	match := filename_pattern.Match([]byte(random_file_name))

	if(match == false) {
		t.Errorf("Random file name generator works incorrectly: %s\n", random_file_name)
	}
}

func TestDownloadFile(t *testing.T) {
	url := "https://raw.githubusercontent.com/ksukhorukov/Atlant/master/samples/sample.csv"
	tmp_file_path := RandomFile("../tmp", 64)

	err := DownloadFile(url, tmp_file_path)

	if(err != nil) {
		t.Errorf("Cannot download sample file: %v\n", err)
	}

	DeleteFile(tmp_file_path)
}

func TestDownloadFileWithWrongMimeType(t *testing.T) {
	url := "https://github.com/ksukhorukov/Atlant/raw/master/samples/golang.png"
	tmp_file_path := RandomFile("../tmp", 64)

	err := DownloadFile(url, tmp_file_path)

	if(err == nil) {
		t.Errorf("Function allows to download files with incorrect mime types\n")
	}

	DeleteFile(tmp_file_path)	
}

func TestSuccessfullySaveResults(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	product := "test_product_1111111111"
	price := 99.9

	result := SaveResults(collection, mng_context, product, price, time.Now().Unix())
		
	if result == false {
		t.Errorf("Cannot save results to MongoDB\n")
	} else {
		_, err := collection.DeleteOne(mng_context, bson.M{"product": product})

		if err != nil {
			t.Errorf("Cannot delete record from MongoDB\n")
		}
	}
}

func TestDontSaveProductsWithTheSamePrice(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	product := "test_product_1111111111"
	price := 99.9

	result := SaveResults(collection, mng_context, product, price, time.Now().Unix())
		
	if result == false {
		t.Errorf("Cannot save results to MongoDB\n")
	} 

	result = SaveResults(collection, mng_context, product, price, time.Now().Unix())	

	if result == true {
		t.Errorf("Can save record with equal prices")
	}

	_, err := collection.DeleteOne(mng_context, bson.M{"product": product})

	if err != nil {
		t.Errorf("Cannot delete record from MongoDB\n")
	}
}

func SaveResultsStub(collection mongo.Collection, mng_context context.Context, product string, price float64, timestamp int64) bool {
	return true
}

func TestParseCSV(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	saver := SaveResultsStub

	file_path := "../samples/sample.csv"

	count, err := ParseCSV(file_path, saver, collection, mng_context, time.Now().Unix())

	if err != nil {
		t.Errorf("Parses returned error: %v", err)
	}

	if count != 1000 {
		t.Errorf("ParseCSV did not exported all records")
	}
}

func TestParseCSVProduceErrorWhenCannotOpenCSVFile(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	saver := SaveResultsStub

	file_path := "../samples/sample_abrakadabra.csv"

	_, err := ParseCSV(file_path, saver, collection, mng_context, time.Now().Unix())

	if err == nil {
		t.Errorf("Parser allows to open non-existing files")
	}
}

func TestParseCSVProduceErrorWhenCSVFilesHasIncorrectHeaders(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	saver := SaveResultsStub

	file_path := "../samples/invalid_headers.csv"

	_, err := ParseCSV(file_path, saver, collection, mng_context, time.Now().Unix())

	if err == nil {
		t.Errorf("Parser successfully parsed CSV with invalid headers: %s\n", file_path)
	}
}

func TestParseCSVProduceErrorWhenCSVFilesHasIncorrectStructure(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	saver := SaveResultsStub

	file_path := "../samples/invalid_structure.csv"

	_, err := ParseCSV(file_path, saver, collection, mng_context, time.Now().Unix())

	if err == nil {
		t.Errorf("Parser successfully parsed CSV with invalid structure: %s\n", file_path)
	}
}

func TestParseCSVProduceErrorWhenCSVFilesHasIncorrectValues(t *testing.T) {
	mongo_address = "127.0.0.1"

	mng_context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, collection := InitMongo(mng_context)

	defer client.Disconnect(mng_context)

	saver := SaveResultsStub

	file_path := "../samples/invalid_structure.csv"

	_, err := ParseCSV(file_path, saver, collection, mng_context, time.Now().Unix())

	if err == nil {
		t.Errorf("Parser successfully parsed CSV with invalid values: %s\n", file_path)
	}
}

func TestGetCursorIndex(t *testing.T) {
	var page int64
	var per_page int64 
	var length int64 
	var expected_start int64 
	var expected_end int64
	var start int64 
	var end int64

	page = 1
	per_page = 10
	length = 20

	expected_start, expected_end = 0, per_page

	start, end = GetCursorRange(page, per_page, length)

	if start != expected_start || end != expected_end {
		t.Errorf("Invalid cursors. Expecting (%d, %d) got (%d, %d)\n", expected_start, expected_end, start, end)
	}

	page = -1
	per_page = 10
	length = 20

	expected_start = 10
	expected_end = 20

	start, end = GetCursorRange(page, per_page, length)

	if start != expected_start || end != expected_end {
		t.Errorf("Invalid cursors. Expecting (%d, %d) got (%d, %d)\n", expected_start, expected_end, start, end)
	}

	page = 2
	per_page = 10
	length = 20

	expected_start = 10
	expected_end = 20

	start, end = GetCursorRange(page, per_page, length)	

	if start != expected_start || end != expected_end {
		t.Errorf("Invalid cursors. Expecting (%d, %d) got (%d, %d)\n", expected_start, expected_end, start, end)
	}

	page = 4
	per_page = 10
	length = 20

	expected_start = 0
	expected_end = per_page

	start, end = GetCursorRange(page, per_page, length)	

	if start != expected_start || end != expected_end {
		t.Errorf("Invalid cursors. Expecting (%d, %d) got (%d, %d)\n", expected_start, expected_end, start, end)
	}


	page = 1
	per_page = 100
	length = 20

	expected_start = 0
	expected_end = length

	start, end = GetCursorRange(page, per_page, length)	

	if start != expected_start || end != expected_end {
		t.Errorf("Invalid cursors. Expecting (%d, %d) got (%d, %d)\n", expected_start, expected_end, start, end)
	}
}

