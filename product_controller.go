package main

import (
	"encoding/json"
	"net/http"
	"versionHistoryService/Application/Godeps/_workspace/src/github.com/gorilla/context"
	"versionHistoryService/Application/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"versionHistoryService/Application/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"versionHistoryService/Application/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"versionHistoryService/Application/Godeps/_workspace/src/gopkg.in/unrolled/render.v1"
)

type ProductController struct {
	AppController
	*render.Render
}

func (c *ProductController) Index(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	var results []Product

	err := db.C("products").Find(nil).All(&results)
	if err != nil {
		return err
	}

	c.JSON(rw, 200, results)

	return nil
}

func (c *ProductController) Create(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	decoder := json.NewDecoder(r.Body)
	var product Product
	var vendor Vendor

	err := decoder.Decode(&product)
	if err != nil {
		return err
	}

	products := db.C("products")
	if err := products.Insert(&Product{Name: product.Name, VendorName: product.VendorName}); err != nil {
		return err
	}

	if err := db.C("vendors").Find(bson.M{"name": product.VendorName}).One(&vendor); err != nil {
		return err
	}

	change := bson.M{"$push": bson.M{"productnames": product.Name}}
	db.C("vendors").Update(vendor, change)

	return err
}

func (c *ProductController) Show(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	params := context.Get(r, "params").(httprouter.Params)
	var results Product

	err := db.C("products").Find(bson.M{"vendorname": params.ByName("vendor_name"), "name": params.ByName("name")}).One(&results)
	if err != nil {
		return err
	}

	c.JSON(rw, 200, results)

	return nil
}
