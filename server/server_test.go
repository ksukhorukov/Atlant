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

func TestSocketAddress(t *testing.T) {
	server_address = "127.0.0.1"
	server_port = 55555

	server_socket := fmt.Sprintf("%s:%d", server_address, server_port)

	func_result := SocketAddress()

	if(server_socket != func_result) {
		t.Errorf("%s != %s\n", func_result, server_socket)	
	}
}

