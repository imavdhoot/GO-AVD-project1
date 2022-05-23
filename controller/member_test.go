package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/GO-AVD-project1/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var newMembId = ""
var newMercId2 = ""

func TestGetMember(t *testing.T) {
	testName := "[TestGetMember]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	req, err := http.NewRequest("GET", "/member/-10", nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.GET("/member/:id", GetMember)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusBadRequest, resp.Code)
}

func TestAddMember(t *testing.T) {
	testName := "[TestAddMember]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	jsonStr := []byte(`{"name": "avdhoot1", "address": "150, orchard road 089775"}`)
	req, err := http.NewRequest("POST", "/merchant/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.POST("/merchant/add", AddMerchant)
	r.POST("/member/add", AddMember)
	r.GET("/member/:id", GetMember)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	var respBody map[string]string
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	newMercId2 = respBody["merchantId"]

	jsonStr = []byte(fmt.Sprintf(`{"name": "jack", "email": "jack@gmail.com", "merchantId": "%s"}`, newMercId2))

	req1, _ := http.NewRequest("POST", "/member/add", bytes.NewBuffer(jsonStr))
	resp1 := httptest.NewRecorder()
	r.ServeHTTP(resp1, req1)
	fmt.Printf("%s Response:: %v\n", testName, resp1)
	assertStatus(t, testName, http.StatusOK, resp1.Code)
	var resp1Body AddMemberResp
	json.Unmarshal(resp1.Body.Bytes(), &resp1Body)
	newMembId = strconv.Itoa(resp1Body.MemberID)

	req2, _ := http.NewRequest("GET", "/member/"+newMembId, nil)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)
	fmt.Printf("%s Response:: %v\n", testName, resp2)
	assertStatus(t, testName, http.StatusOK, resp2.Code)

	var resp2Body map[string]string
	json.Unmarshal(resp2.Body.Bytes(), &resp2Body)
	assertStringValues(t, testName, "jack", resp2Body["name"])

}

func TestUpdateMember(t *testing.T) {
	testName := "[TestUpdateMember]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	jsonStr := []byte(`{"email": "jack100@gmail.com"}`)
	req, err := http.NewRequest("PUT", "/member/"+newMembId, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.PUT("/member/:id", UpdateMember)
	r.GET("/member/:id", GetMember)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	req2, _ := http.NewRequest("GET", "/member/"+newMembId, nil)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)
	fmt.Printf("%s Response:: %v\n", testName, resp2)
	assertStatus(t, testName, http.StatusOK, resp2.Code)

	var resp2Body map[string]string
	json.Unmarshal(resp2.Body.Bytes(), &resp2Body)
	assertStringValues(t, testName, "jack100@gmail.com", resp2Body["email"])
}

func TestMemberListByMerchant(t *testing.T) {
	testName := "[TestMemberListByMerchant]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	jsonStr := []byte(fmt.Sprintf(`{"name": "jack2", "email": "jack2@gmail.com", "merchantId": "%s"}`, newMercId2))
	req, err := http.NewRequest("POST", "/member/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.POST("/member/add", AddMember)
	r.GET("/members/list/:merchantId", MemberListByMerchant)
	r.DELETE("/member/:id", DeleteMember)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	var respBody AddMemberResp
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	newMembId2 := strconv.Itoa(respBody.MemberID)

	req1, _ := http.NewRequest("GET", "/members/list/"+newMercId2, nil)
	resp1 := httptest.NewRecorder()
	r.ServeHTTP(resp1, req1)
	fmt.Printf("%s Response:: %v\n", testName, resp1)
	assertStatus(t, testName, http.StatusOK, resp1.Code)
	var resp1Body MemberListByMerchantResp
	json.Unmarshal(resp1.Body.Bytes(), &resp1Body)

	assertStringValues(t, testName, "members fetched successfully", resp1Body.Message)

	req2, _ := http.NewRequest("DELETE", "/member/"+newMembId2, nil)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)
	fmt.Printf("%s Response:: %v\n", testName, resp2)
	assertStatus(t, testName, http.StatusOK, resp2.Code)
}

func TestDeleteMember(t *testing.T) {
	testName := "[TestDeleteMember]"
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	model.InitDB()

	// Create new HTTP request
	//json := []byte(``)
	req, err := http.NewRequest("DELETE", "/member/"+newMembId, nil)
	if err != nil {
		t.Fatalf("%s Couldn't create request: %v\n", testName, err)
	}

	//Assign HTTP handler fn
	r.DELETE("/member/:id", DeleteMember)
	r.DELETE("/merchant/:id", DeleteMerchant)
	r.GET("/member/:id", GetMember)

	//create recorder for HTTP response
	resp := httptest.NewRecorder()

	//fire the request and receive the response
	r.ServeHTTP(resp, req)

	// Add required assertions on response
	fmt.Printf("%s Response:: %v\n", testName, resp)
	assertStatus(t, testName, http.StatusOK, resp.Code)

	req2, _ := http.NewRequest("GET", "/member/"+newMembId, nil)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)
	fmt.Printf("%s Response:: %v\n", testName, resp2)
	assertStatus(t, testName, http.StatusBadRequest, resp2.Code)

	req1, _ := http.NewRequest("DELETE", "/merchant/"+newMercId2, nil)
	resp1 := httptest.NewRecorder()
	r.ServeHTTP(resp1, req1)
	fmt.Printf("%s Response:: %v\n", testName, resp1)
	assertStatus(t, testName, http.StatusOK, resp1.Code)

}
