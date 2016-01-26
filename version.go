package main

import (
	"time"
)

type Version struct {
        ReleaseNumber string `json:"releasenumber"`
        ReleaseDate time.Time `json:"releasedate"`
        ProductName string `json:"productname"`
        VendorName string `json:"vendorname"`
}