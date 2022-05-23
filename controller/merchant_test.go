package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, testName string, expected, actual int) {
	if expected != actual {
		t.Errorf(" %s Expected response code %d. Got %d\n", testName, expected, actual)
	}
}

func TestGetMerchant(t *testing.T) {
	testName := "[GetMerchant] Fetch"
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	// Create new HTTP request
	//json := []byte(``)
	req, err := http.NewRequest("GET", "localhost:9091/merchant/merc1", nil)
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
	fmt.Println(resp.Body)
	checkResponseCode(t, testName, http.StatusNotFound, resp.Code)
}
