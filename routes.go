package main

import (
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"
    "gopkg.in/unrolled/render.v1"
)

func Router(db *mgo.Database) *httprouter.Router {
	r := httprouter.New()

	verndor_controller := &VendorController{Render: render.New(render.Options{})}

	// GET /vendors
	r.GET("/vendors", verndor_controller.Action(verndor_controller.Index, db))

	// GET /vendors/name
	r.GET("/vendors/:name", verndor_controller.Action(verndor_controller.Show, db))
	
	// POST /vendors
	r.POST("/vendors", verndor_controller.Action(verndor_controller.Create, db))


	product_controller := &ProductController{Render: render.New(render.Options{})}

	// GET /products
	r.GET("/products", product_controller.Action(product_controller.Index, db))

	// GET /products/vendor_name/name
	r.GET("/products/:vendor_name/:name", product_controller.Action(product_controller.Show, db))
	
	// POST /products
	r.POST("/products", product_controller.Action(product_controller.Create, db))


	version_controller := &VersionController{Render: render.New(render.Options{})}

	// GET /versions
	r.GET("/versions", version_controller.Action(version_controller.Index, db))

	// GET /versions/vendor_name/product_name/release_number
	r.GET("/versions/:vendor_name/:product_name/:release_number", version_controller.Action(version_controller.Show, db))
	
	// POST /versions
	r.POST("/versions", version_controller.Action(version_controller.Create, db))


	return r
}