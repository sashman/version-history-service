package main

type Product struct {
        Name string `json:"name"`
        VendorName string `json:"vendorname"`
        LatestVersion Version `json:"latestversion"`
}