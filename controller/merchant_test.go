package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/GO-AVD-project1/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var newMercId = ""

func TestGetMerchant(t *testing.T) {
	testName := "[TestGetMerchant]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	req, err := http.NewRequest("GET", "/merchant/avdhoot", nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.GET("/merchant/:id", GetMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusBadRequest, resp.Code)
}

func TestAddMerchant(t *testing.T) {
	testName := "[TestAddMerchant]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	jsonStr := []byte(`{"name": "avdhoot", "address": "150, orchard road 089775"}`)
	req, err := http.NewRequest("POST", "/merchant/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.POST("/merchant/add", AddMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	var respBody map[string]string
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	newMercId = respBody["merchantId"]
}

func TestGetMerchant2(t *testing.T) {
	testName := "[TestGetMerchant2]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	//json := []byte(``)
	req, err := http.NewRequest("GET", "/merchant/"+newMercId, nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.GET("/merchant/:id", GetMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)
}

func TestUpdateMerchant(t *testing.T) {
	testName := "[TestUpdateMerchant]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	jsonStr := []byte(`{"address": "250, orchard road 089775"}`)
	req, err := http.NewRequest("PUT", "/merchant/"+newMercId, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.PUT("/merchant/:id", UpdateMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)
}

func TestGetMerchant3(t *testing.T) {
	testName := "[TestGetMerchant3]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	req, err := http.NewRequest("GET", "/merchant/"+newMercId, nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.GET("/merchant/:id", GetMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	var respBody map[string]string
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	newAddress := respBody["address"]

	assertStringValues(t, testName, "250, orchard road 089775", newAddress)
}

func TestDeleteMerchant(t *testing.T) {
	testName := "[TestDeleteMerchant]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	//json := []byte(``)
	req, err := http.NewRequest("DELETE", "/merchant/"+newMercId, nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.DELETE("/merchant/:id", DeleteMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)
}

func TestGetMerchant4(t *testing.T) {
	testName := "[TestGetMerchant4]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	req, err := http.NewRequest("GET", "/merchant/"+newMercId, nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.GET("/merchant/:id", GetMerchant)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusBadRequest, resp.Code)
}
