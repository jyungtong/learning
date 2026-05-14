package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "net/http/httputil"
	"os"
)

var TG_BASE_URL = "https://api.telegram.org"
var TG_TOKEN = os.Getenv("TG_TOKEN")

type User struct {
	Id int
	Name string
}

func main() {
	// if TG_TOKEN == "" {
	// 	log.Fatalf("TG_TOKEN is not set")
	// 	return
	// }

	// resp, err := http.Get(fmt.Sprintf("%s/%s/getUpdates", TG_BASE_URL, TG_TOKEN))
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Println("getUpdates err:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("resp.Body err:", err)
	}
	//
	// fmt.Println(body)
	// respBody, _ := httputil.DumpResponse(resp, true)
	// fmt.Printf("%s", respBody)

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatalf("unable decode json: %s", err)
	}

	fmt.Println("ID:", user.Id, "name:", user.Name)
}
