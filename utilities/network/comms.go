package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

import (
	configs "github.com/acheong08/SimpleResv-Client/configs"
	data "github.com/acheong08/SimpleResv-Client/data"
)

var BaseURL string

func init() {
	BaseURL = configs.ServerProtocol + configs.ServerHost + configs.ServerPort
}
func postData(reqURL string, request data.Request) string {
	// Convert data to json
	jsonValues, _ := json.Marshal(request)
	// Post data
	resp, err := http.Post(reqURL, "application/json", bytes.NewBuffer(jsonValues))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

////////////////////////////////////////////// User ////////////////////////////////////////////////////////////
// Get items
func GetItems() []data.Item {
	// Build URL
	var reqURL string = BaseURL + "/api/GetItems"
	// Get data from server
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rawbody, err := ioutil.ReadAll(resp.Body)
	body := string(rawbody)
	if err != nil {
		log.Fatal(err)
	}
	// Convert JSON data to struct
	var items []data.Item
	err = json.Unmarshal([]byte(body), &items)
	if err != nil {
		log.Fatal(err)
	}
	return items
}

// Auth
func AuthUser(email string, password string) bool {
	// Set request URL
	var reqURL string = BaseURL + "/api/CheckAuth"
	// Set data
	var request data.Request
	request.Email = email
	request.Password = password
	// Post data
	body := postData(reqURL, request)
	// Check authorization
	if body == "true" {
		return true
	} else {
		return false
	}
}

// Borrow
func TakeItem(email string, password string, id int) string {
	var reqURL string = BaseURL + "/api/User"
	var request data.Request
	request.Action = "ToggleItem"
	request.Email = email
	request.Password = password
	request.Id = id
	request.Available = false
	request.Status = email
	// Return output
	return postData(reqURL, request)
}

// Return
func ReturnItem(email string, password string, id int) string {
	var reqURL string = BaseURL + "/api/User"
	var request data.Request
	request.Action = "ToggleItem"
	request.Email = email
	request.Password = password
	request.Id = id
	request.Available = true
	request.Status = "Available"
	return postData(reqURL, request)
}

////////////////////////////////////////////// Admin ////////////////////////////////////////////////////////////
// Add user
func AddUser(email string, password string, addemail string, addpassword string) string {
	var reqURL string = BaseURL + "/api/Admin"
	var request data.Request
	request.Action = "AddUser"
	request.Email = email
	request.Password = password
	request.AddEmail = addemail
	request.AddPassword = addpassword
	return postData(reqURL, request)
}

// Delete user
func DeleteUser(email string, password string, addemail string) string {
	var reqURL string = BaseURL + "/api/Admin"
	var request data.Request
	request.Action = "DeleteUser"
	request.Email = email
	request.Password = password
	request.AddEmail = addemail
	return postData(reqURL, request)
}

// Add item
func AddItem(email string, password string, name string, details string) string {
	var reqURL string = BaseURL + "/api/Admin"
	var request data.Request
	request.Action = "AddItem"
	request.Email = email
	request.Password = password
	request.Name = name
	request.Details = details
	return postData(reqURL, request)
}

// Delete item
func DeleteItem(email string, password string, id int) string {
	var reqURL string = BaseURL + "/api/Admin"
	var request data.Request
	request.Action = "DeleteItem"
	request.Email = email
	request.Password = password
	request.Id = id
	return postData(reqURL, request)
}

// Reset Database
func ResetDB(email string, password string) string {
	var reqURL string = BaseURL + "/api/Admin"
	var request data.Request
	request.Action = "Reset"
	request.Email = email
	request.Password = password
	return postData(reqURL, request)
}
