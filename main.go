package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Api_result struct {
	Url             string `json:"url"`
	Date            string `json:"date"`
	HDurl           string `json:"hdurl"`
	Title           string `json:"title"`
	Copyright       string `json:"copyright"`
	Media_type      string `json:"media_type"`
	Explanation     string `json:"explanation"`
	Service_version string `json:"service_version"`
}

var api_key string = "your_nasa_api_key_here"

func main() {

	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&hd=true", api_key)

	req, err := http.NewRequest("GET", url, nil)
	check_error(err, "NewRequest : ")

	client := &http.Client{}

	resp, err := client.Do(req)
	check_error(err, "Do: ")

	defer resp.Body.Close()

	var data Api_result

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)
	}

	fmt.Println(data)

}

func check_error(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		return
	}
}
