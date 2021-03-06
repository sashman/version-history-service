package main

import (
	"encoding/json"
	"net/http"
	"versionHistoryService/Godeps/_workspace/src/github.com/gorilla/context"
	"versionHistoryService/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/unrolled/render.v1"
)

type VersionController struct {
	AppController
	*render.Render
}

func (c *VersionController) Index(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	var results []Version

	err := db.C("versions").Find(nil).All(&results)
	if err != nil {
		return err
	}

	c.JSON(rw, 200, results)

	return nil
}

func (c *VersionController) Create(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	decoder := json.NewDecoder(r.Body)
	var version Version
	var product Product
	var vendor Vendor

	err := decoder.Decode(&version)
	if err != nil {
		return err
	}

	versions := db.C("versions")

	index := mgo.Index{
		Key:    []string{"releasedate", "releasenumber", "productname", "vendorname"},
		Unique: true,
	}
	err = versions.EnsureIndex(index)
	if err != nil {
		return err
	}

	if err := versions.Insert(&Version{version.ReleaseNumber, version.ReleaseDate, version.ProductName, version.VendorName}); err != nil {
		return err
	}

	if err := db.C("products").Find(bson.M{"name": version.ProductName, "vendorname": version.VendorName}).One(&product); err != nil {
		if err == mgo.ErrNotFound {

			products := db.C("products")
			if err := products.Insert(&Product{Name: version.ProductName, VendorName: version.VendorName}); err != nil {
				return err
			}

			if err := db.C("vendors").Find(bson.M{"name": version.VendorName}).One(&vendor); err != nil {
				if err == mgo.ErrNotFound {
					verndors := db.C("vendors")

					if err := verndors.Insert(&Vendor{Name: version.VendorName}); err != nil {
						return err
					}

					if err := db.C("vendors").Find(bson.M{"name": version.VendorName}).One(&vendor); err != nil {
						return err
					}

				} else {
					return err
				}
			}

			change := bson.M{"$push": bson.M{"productnames": product.Name}}
			db.C("vendors").Update(vendor, change)

			if err := db.C("products").Find(bson.M{"name": version.ProductName, "vendorname": version.VendorName}).One(&product); err != nil {
				return err
			}

		} else {
			return err
		}
	}

	if &product.LatestVersion == nil || product.LatestVersion.ReleaseDate.Before(version.ReleaseDate) {
		change := bson.M{"$set": bson.M{"latestversion": version}}
		if err := db.C("products").Update(bson.M{"name": version.ProductName, "vendorname": version.VendorName}, change); err != nil {
			return err
		}
	}

	return err
}

func (c *VersionController) Show(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error {
	params := context.Get(r, "params").(httprouter.Params)
	var results Version

	err := db.C("versions").Find(bson.M{"vendorname": params.ByName("vendor_name"), "productname": params.ByName("name"), "releasenumber": params.ByName("release_number")}).One(&results)
	if err != nil {
		return err
	}

	c.JSON(rw, 200, results)

	return nil
}
