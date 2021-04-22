package main

import (
	"testing"
	"fmt"
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