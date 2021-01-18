package main

import (
	"fmt"
	"net/http"

	// "io/ioutil"
	"io"
	"os"
)

// TODO: Add input for arguements
func main() {
	// input for site
	fmt.Println("Enter Website to scrape:")
	var site string
	fmt.Scanln(&site)

	//Make GET request
	f, _ := os.Create("./test.html")
	resp, err := http.Get(site)
	Check(err)

	//close the connection stream
	defer resp.Body.Close()

	//get the bytes acquired
	//resultsInBytes,err := ioutil.ReadAll(resp.Body)
	resultsInBytes, err := io.Copy(f, resp.Body)
	Check(err)

	//convert into a readable format
	results := string(resultsInBytes)

	//Display website
	fmt.Println(results)
}

//error handler function
func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
