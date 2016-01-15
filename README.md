# versionHistoryService
Web API to store versions and release based on products and vendors. This is primary aimed at storing antivirus version history

## Prerequisites
* Go https://golang.org/dl/
* MongoDb https://www.mongodb.org/downloads#production

## Set up
* Start Mongo with `mongod`
* Build project `go build`
* Run `versionHistoryService`
* The service should be running on `localhost:8080` by default

## Available endpoints
* GET /vendors
* GET /vendors/name
* POST /vendors
  * `{ "name": "test" }`
* GET /products
* GET /products/vendor_name/name
* POST /products
  * `{ "name": "test_product", "vendorname": "" }`
* GET /versions
* GET /versions/vendor_name/product_name/release_number
* POST /versions
  * `{ "releasenumber": "1.2.123", "releasedate": "0001-01-01T00:00:00Z", "productname": "test_product", "vendorname": "test" }`
