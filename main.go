package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const MyGetUrl = "http://localhost:6000/get"
const MyPostUrl = "http://localhost:6000/post"
const MyPostFormUrl = "http://localhost:6000/postform"

func main() {
	fmt.Println("Welcome to Requests")
	//GET Request :
	//PerformGetRequest(MyUrl)
	//POST Request :
	// PerformPostJsonRequest(MyPostUrl)
	//POST Form Request :
	PerformPostFormRequest(MyPostFormUrl)
}

// Here P below for performgetrequest means it is made PUBLIC
func PerformGetRequest(MyGetUrl string) {

	response, err := http.Get(MyGetUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("The Response Data from the GET Request is : \n", string(content))
}

func PerformPostJsonRequest(MyPostUrl string) {
	// fake json payload/data :
	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go",
		"price":0,
		"platform":"loc"
	}
`)

	response, err := http.Post(MyPostUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println("Data from POST request is : \n", string(content))

}

func PerformPostFormRequest(MyPostFormUrl string) {
	// For sending form-data :
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh.go@dev")

	response, err := http.PostForm(MyPostFormUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("The Response Data from Form POST request is : \n", string(content))

}
