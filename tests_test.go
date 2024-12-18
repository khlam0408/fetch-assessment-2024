package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fetch-assessment/models"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	//"github.com/gin-gonic/gin"
	"strings"
	//"fmt"
	"github.com/google/uuid"

)

// Testing for one receipt with all types of rules that can be added 

func TestExampleReceiptAllTypeRules(t *testing.T){
	exampleReceipt := models.Receipt{ 
		Retailer: "M&M Corner Market", // Test Alpha Numeric characters and non-alpha numeric characters
  		PurchaseDate: "2022-03-21", // Test Odd days
  		PurchaseTime: "14:33", // Test Between 2:00pm - 4:00pm
		Total: "37.00", // Test Multiple of 0.25 and round dollar amount with no cents
		Items: []map[string]string{ // Tests 2 items per pair, (so here it is only 4 pairs since theres 5 items)
			{
				"shortDescription": "Gatorade",
				"price": "6.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "4.25",
			},
			{
				"shortDescription": "Emils Cheese Pizza", // Tests if item's descript length is mult of 3
				"price": "12.25",
			},
			{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", // Tests of item's descript length is mult of 3 as well as trimmed length
				"price": "12.00",
			},
		},
	}

	r := SetUpRouter()
	w1 := httptest.NewRecorder()
	receiptJson, _ := json.Marshal(exampleReceipt)
	reqPost, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
	r.ServeHTTP(w1, reqPost)
	var id1 models.ID
	//responseW, _ := json.Marshal(w)
	errPost := json.Unmarshal([]byte(w1.Body.String()), &id1)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id1)
	assert.Nil(t, errPost)

	w2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "/receipts/" + id1.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet)
	expectedResponse := map[string]int{
		"points": 121,
	}
	var actualResponse map[string]int
	errGet := json.Unmarshal([]byte(w2.Body.String()), &actualResponse)
	assert.Nil(t, errGet)
	assert.Equal(t, expectedResponse, actualResponse)

}

// Testing one receipt with a few rules (not all type of rules)

func TestExampleReceiptFewRules1(t *testing.T){
	exampleReceipt := models.Receipt{ 
		Retailer: "Target", // Test Alpha Numeric characters
  		PurchaseDate: "2022-01-01", // Tests for odd day
  		PurchaseTime: "13:01",
		Total: "35.35",
		Items: []map[string]string{ // Tests the pair of 2 items, only 2 groups
			{
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49",
			},
			{
				"shortDescription": "Emils Cheese Pizza", // Tests if length of descript is mult of 3
				"price": "12.25",
			},
			{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26",
			},
			{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35",
			},
			{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", //Tests if length of descript is mult of 3 and trimmed length
				"price": "12.00",
			},
		},
	}

	r := SetUpRouter()
	w1 := httptest.NewRecorder()
	receiptJson, _ := json.Marshal(exampleReceipt)
	reqPost, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
	r.ServeHTTP(w1, reqPost)
	var id1 models.ID
	//responseW, _ := json.Marshal(w)
	errPost := json.Unmarshal([]byte(w1.Body.String()), &id1)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id1)
	assert.Nil(t, errPost)

	w2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "/receipts/" + id1.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet)
	expectedResponse := map[string]int{
		"points": 28,
	}
	var actualResponse map[string]int
	errGet := json.Unmarshal([]byte(w2.Body.String()), &actualResponse)
	assert.Nil(t, errGet)
	assert.Equal(t, expectedResponse, actualResponse)

}

// Testing a receipt with few rules (not all types of rules)

func TestExampleReceiptFewRules2(t *testing.T){
	exampleReceipt := models.Receipt{ 
		Retailer: "M&M Corner Market", // Tests none alphanumeric characters 
  		PurchaseDate: "2022-03-20",
  		PurchaseTime: "14:33", // Tests if time is between 2pm and 4pm
		Total: "9.00", // Tests total is mult of .25 as well as rounded cents 
		Items: []map[string]string{ // Tests pairs of items 
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
		},
	}
	

	r := SetUpRouter()
	w1 := httptest.NewRecorder()
	receiptJson, _ := json.Marshal(exampleReceipt)
	reqPost, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson)))
	r.ServeHTTP(w1, reqPost)
	var id1 models.ID
	//responseW, _ := json.Marshal(w)
	errPost := json.Unmarshal([]byte(w1.Body.String()), &id1)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id1)
	assert.Nil(t, errPost)

	w2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "/receipts/" + id1.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet)
	expectedResponse := map[string]int{
		"points": 109,
	}
	var actualResponse map[string]int
	errGet := json.Unmarshal([]byte(w2.Body.String()), &actualResponse)
	assert.Nil(t, errGet)
	assert.Equal(t, expectedResponse, actualResponse)

}

// Tests multiple receipts and checks if all the points are adding up correctly and in the correct ids 

func TestExampleReceiptMultiple(t *testing.T){
	// This whole tests if code can handle multiple receipts and not mix the receipts up via the ids. 
	exampleReceipt1 := models.Receipt{ 
		Retailer: "M&M Corner Market",
  		PurchaseDate: "2022-03-20",
  		PurchaseTime: "14:33",
		Total: "9.00",
		Items: []map[string]string{
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
			{
				"shortDescription": "Gatorade",
				"price": "2.25",
			},
		},
	}
	exampleReceipt2 := models.Receipt{ 
		Retailer: "Target",
  		PurchaseDate: "2022-01-01",
  		PurchaseTime: "13:01",
		Total: "35.35",
		Items: []map[string]string{
			{
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49",
			},
			{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25",
			},
			{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26",
			},
			{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35",
			},
			{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00",
			},
		},
	}

	exampleReceipt3 := models.Receipt{ 
		Retailer: "Walgreens",
  		PurchaseDate: "2022-01-02",
  		PurchaseTime: "08:13",
		Total: "2.65",
		Items: []map[string]string{
			{
				"shortDescription": "Pepsi - 12-oz",
				"price": "1.25",
			},
			{
				"shortDescription": "Dasani",
				"price": "1.40",
			},
		},
	}

	exampleReceipt4 := models.Receipt{ 
		Retailer: "Target",
  		PurchaseDate: "2022-01-02",
  		PurchaseTime: "13:13",
		Total: "1.25",
		Items: []map[string]string{
			{
				"shortDescription": "Pepsi - 12-oz",
				"price": "1.25",
			},
		},
	}

	r := SetUpRouter()
	w1 := httptest.NewRecorder()
	receiptJson1, _ := json.Marshal(exampleReceipt1)
	reqPost1, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson1)))
	r.ServeHTTP(w1, reqPost1)
	var id1 models.ID
	//responseW, _ := json.Marshal(w)
	errPost1 := json.Unmarshal([]byte(w1.Body.String()), &id1)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id1)
	assert.Nil(t, errPost1)

	/////////

	w1 = httptest.NewRecorder()
	receiptJson2, _ := json.Marshal(exampleReceipt2)
	reqPost2, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson2)))
	r.ServeHTTP(w1, reqPost2)
	var id2 models.ID
	//responseW, _ := json.Marshal(w)
	errPost2 := json.Unmarshal([]byte(w1.Body.String()), &id2)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id2)
	assert.Nil(t, errPost2)

	//////////

	w1 = httptest.NewRecorder()
	receiptJson3, _ := json.Marshal(exampleReceipt3)
	reqPost3, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson3)))
	r.ServeHTTP(w1, reqPost3)
	var id3 models.ID
	//responseW, _ := json.Marshal(w)
	errPost3 := json.Unmarshal([]byte(w1.Body.String()), &id3)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id3)
	assert.Nil(t, errPost3)

	///////////

	w1 = httptest.NewRecorder()
	receiptJson4, _ := json.Marshal(exampleReceipt4)
	reqPost4, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(string(receiptJson4)))
	r.ServeHTTP(w1, reqPost4)
	var id4 models.ID
	//responseW, _ := json.Marshal(w)
	errPost4 := json.Unmarshal([]byte(w1.Body.String()), &id4)
	
	assert.Equal(t, 201, w1.Code)
	assert.NotEqual(t, "", id4)
	assert.Nil(t, errPost4)
	
	////////////////////////////////////////////////////////////////////////////////////////////////////////////

	w2 := httptest.NewRecorder()
	reqGet1, _ := http.NewRequest("GET", "/receipts/" + id1.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet1)
	expectedResponse1 := map[string]int{
		"points": 109,
	}
	var actualResponse1 map[string]int
	errGet1 := json.Unmarshal([]byte(w2.Body.String()), &actualResponse1)
	assert.Equal(t, 200, w2.Code)
	assert.Nil(t, errGet1)
	assert.Equal(t, expectedResponse1, actualResponse1)

	//////////

	w2 = httptest.NewRecorder()
	reqGet2, _ := http.NewRequest("GET", "/receipts/" + id2.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet2)
	expectedResponse2 := map[string]int{
		"points": 28,
	}
	var actualResponse2 map[string]int
	errGet2 := json.Unmarshal([]byte(w2.Body.String()), &actualResponse2)
	assert.Equal(t, 200, w2.Code)
	assert.Nil(t, errGet2)
	assert.Equal(t, expectedResponse2, actualResponse2)

	//////////

	w2 = httptest.NewRecorder()
	reqGet3, _ := http.NewRequest("GET", "/receipts/" + id3.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet3)
	expectedResponse3 := map[string]int{
		"points": 15,
	}
	var actualResponse3 map[string]int
	errGet3 := json.Unmarshal([]byte(w2.Body.String()), &actualResponse3)
	assert.Equal(t, 200, w2.Code)
	assert.Nil(t, errGet3)
	assert.Equal(t, expectedResponse3, actualResponse3)

	///////////

	w2 = httptest.NewRecorder()
	reqGet4, _ := http.NewRequest("GET", "/receipts/" + id4.ID + "/points", nil)
	r.ServeHTTP(w2, reqGet4)
	expectedResponse4 := map[string]int{
		"points": 31,
	}
	var actualResponse4 map[string]int
	errGet4 := json.Unmarshal([]byte(w2.Body.String()), &actualResponse4)
	assert.Equal(t, 200, w2.Code)
	assert.Nil(t, errGet4)
	assert.Equal(t, expectedResponse4, actualResponse4)

}

func TestExampleErrorReturned(t *testing.T){
	// Tests if we input a random id, if it returns an error as the id was never put into the code as a receipt first. 
	newID := uuid.New()

	r := SetUpRouter()
	w := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "/receipts/" + newID.String() + "/points", nil)
	r.ServeHTTP(w, reqGet)
	expectedResponse := map[string]string{
		"message": "id not found",
	}
	var actualResponse map[string]string
	errGet := json.Unmarshal([]byte(w.Body.String()), &actualResponse)
	assert.Nil(t, errGet)
	assert.Equal(t, expectedResponse, actualResponse)
}


