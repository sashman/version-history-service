package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
    "github.com/gorilla/context"
)

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(rw http.ResponseWriter, r *http.Request, db *mgo.Database) error

// This is our Base Controller
type AppController struct{}


// The action function helps with error handling in a controller
func (c *AppController) Action(a Action, db *mgo.Database) httprouter.Handle {
    return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    	context.Set(r, "params", p)
        if err := a(rw, r, db); err != nil {
            http.Error(rw, err.Error(), 500)
        }
    }
}