package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	basicAuth()
}

func basicAuth() string {
	var username string = "dmadmin"
	var passwd string = "password"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.2.132:8080/dctm-rest/repositories/MyRepo/currentuser.json", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
	return s
}
