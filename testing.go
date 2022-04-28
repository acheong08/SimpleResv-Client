package main

import (
  network "github.com/acheong08/SimpleResv-Client/utilities/network"
  gui "github.com/acheong08/SimpleResv-Client/utilities/GUI"
  _ "github.com/acheong08/SimpleResv-Client/data"
  _ "encoding/json"
  "fmt"
)

// Set auth
var email string = "admin@example.com"
var password string = "ExamplePassword"

func main()  {
  // Reset database
  fmt.Println(network.ResetDB(email, password))
  // Success
  //testSuccess()
  // Fail
  //testFail()
  // GUI
  gui.Run()
}

func testSuccess()  {
  fmt.Println("STARTING SUCCESS TEST")
  // Add item
  fmt.Println(network.AddItem(email, password, "Item 1", "Description 1"))
  fmt.Println(network.AddItem(email, password, "Item 2", "Description 2"))
  fmt.Println(network.AddItem(email, password, "Item 3", "Description 3"))
  // Delete items
  fmt.Println(network.DeleteItem(email, password, 1))
  // Add user
  fmt.Println(network.AddUser(email, password, "user@example.com", "UserPassword"))
  // Authenticate user
  fmt.Println(network.AuthUser("user@example.com", "UserPassword"))
  // Delete user
  fmt.Println(network.DeleteUser(email, password, "user@example.com"))
}

func testFail()  {
  fmt.Println("STARTING FAIL TEST")
  // Failed authentication
  fmt.Println(network.AuthUser("user@example.com", "UserPassword"))
  // Add existing item (Should fail)
  fmt.Println(network.AddItem(email, password, "Item 2", "Description 2"))
  // Accessing admin API with user password
  network.AddUser(email, password, "user@example.com", "UserPassword")
  fmt.Println(network.DeleteUser("user@example.com", "UserPassword", email))
}
