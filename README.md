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

#### View all vendors
* GET /vendors

#### View specific vendor
* GET /vendors/name

#### Create a new vendor
* POST /vendors
  * Example post body: `{ "name": "test" }`

#### View all products
* GET /products

#### View specific product of a vendor
* GET /products/vendor_name/name

#### Create a product
* POST /products
  * Example post body: `{ "name": "test_product", "vendorname": "test" }`

#### View all versions
* GET /versions

#### View a release number for a product of a vendor
* GET /versions/vendor_name/product_name/release_number

#### Create a new version
* POST /versions
  * Example post body: `{ "releasenumber": "1.2.123", "releasedate": "0001-01-01T00:00:00Z", "productname": "test_product", "vendorname": "test" }`
