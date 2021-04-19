package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"os"
)

const ERROR_INCORRECT_HEADERS = "Incorrect CSV file headers"
const ERROR_INCORRECT_STRUCTURE = "Incorrect CSV file structure"
const ERROR_INCORRECT_FILE_TYPE = "Incorrect file type"
const DOWNLOAD_DIRECTORY = "./tmp"

func main() {
	args := os.Args

	file_url := args[1]

	file_path := randomFile(DOWNLOAD_DIRECTORY, 64)

	err := downloadFile(file_url, file_path)

	errorCheck(err)

	fmt.Printf("[+] Starting to parse %s\n", file_path)

	readCSV(file_path)

	err = deleteFile(file_path)

	errorCheck(err)
}

func readCSV(file_path string) {
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
	}
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
