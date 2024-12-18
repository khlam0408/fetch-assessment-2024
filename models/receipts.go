package models

type Receipt struct{
	Retailer string
	PurchaseDate string
	PurchaseTime string
	Total string
	Items []map[string] string
}