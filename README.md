# Atlant

## Description

Golang gRPC server and a client inside Docker containers  under HAProxy.
MongoDB used for data storage.

There are two gRPC methods:

- Fetch(URL) - parses external CSV file with the following format: PRODUCT NAME;PRICE.

The last price of each product is saved in DB collection with the timestamp and number of revisions.
 
- List(paging params, sorting params) - Get the list of products according to filtering criterias.

## Install && Deploy

Make sure protobuf installed.

Make sure that local instances of MongoDB and HAproxy turned off.

``make compile``

``docker-compose up -d``

``make test``

HAProxy binds port 5555 to host and at this stage it should be available for connections.

## Usage

``$./server/server --help``

``$ ./client/client --help``

## Helpers

Generator is a simple Ruby script that creates CSV files with random content for testing purposes.

``./generator/generator.rb sample_file_name.csv`` 

## Examples

Bind server to all interfaces and use port 5555,  use custom address for Mongo:

``./server/server --host=0.0.0.0 --port=55555 --mongo_address=192.168.0.100``

Connect to server using socket address and fetch CSV file from local Rails server:

``./client/client --server=localhost:5555 --url=http://localhost:3000/products.csv``
