package main

import (
    "net/http"
    "gopkg.in/mgo.v2"
    "encoding/json"
    "gopkg.in/unrolled/render.v1"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2/bson"
    "github.com/julienschmidt/httprouter"
)

type VendorController struct {
    AppController
    *render.Render
}

func (c *VendorController) Index(rw http.ResponseWriter, r *http.Request, db *mgo.Database ) error {
    var results []Vendor

    err := db.C("vendors").Find(nil).All(&results)
    if err != nil {
        return err
    }

    c.JSON(rw, 200, results)

    return nil
}

func (c *VendorController) Create(rw http.ResponseWriter, r *http.Request, db *mgo.Database ) error {
    decoder := json.NewDecoder(r.Body)
    var vendor Vendor   

    err := decoder.Decode(&vendor)
    if err != nil {
        return err    
    }
    
    verndors := db.C("vendors")
    err = verndors.Insert(&Vendor{vendor.Name, vendor.ProductNames})

    return err
}

func (c *VendorController) Show(rw http.ResponseWriter, r *http.Request, db *mgo.Database ) error {
    params := context.Get(r, "params").(httprouter.Params)
    var results Vendor

    err := db.C("vendors").Find(bson.M{"name": params.ByName("name")}).One(&results)
    if err != nil {
        return err
    }

    c.JSON(rw, 200, results)

    return nil
}