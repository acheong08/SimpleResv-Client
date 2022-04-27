package main

import (
  network "github.com/acheong08/SimpleResv-Client/utilities/network"
  data "github.com/acheong08/SimpleResv-Client/data"
  "encoding/json"
  "fmt"
)

func main()  {
  // Set auth
  var email string = "admin@example.com"
  var password string = "ExamplePassword"
  // Add item
  fmt.Println(network.AddItem(email, password, "Item 1"))
  fmt.Println(network.AddItem(email, password, "Item 2"))
  // Get items
  var items []data.Item = network.GetItems()
  itemjson, _ := json.Marshal(items)
  fmt.Println(string(itemjson))
}
