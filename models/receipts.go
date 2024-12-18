package models

// Structure for the Receipts, used in the testing files
type Receipt struct{
	Retailer string
	PurchaseDate string
	PurchaseTime string
	Total string
	Items []map[string] string
}