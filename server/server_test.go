package main

import (
	"github.com/gabriel-vasile/mimetype"
	"testing"
	"regexp"
	"io/ioutil"
	"fmt"
	"os"
)

func TestGetCursorIndex(t *testing.T) {
	fmt.Println("GetCursorIndex test")
}

func TestMongoAddress(t *testing.T) {
	mongo_address = "127.0.0.1"
	mongo_port = 27017

	mongo_url := fmt.Sprintf("mongodb://%s:%d", mongo_address, mongo_port)
	
	func_result := MongoAddress()

	if(func_result != mongo_url) {
		t.Errorf("%s != %s\n", func_result, mongo_url)
	}
}

func TestSocketAddress(t *testing.T) {
	server_address = "127.0.0.1"
	server_port = 55555

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