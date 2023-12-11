package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseUrl = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error !", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID : %s\t Name :%s\t Grade : %d\n", each.ID, each.Name, each.Grade)
	}

	var user, errr = fetchUser("E001")
	if errr != nil {
		fmt.Println("Error !", errr.Error())
		return
	}
	fmt.Printf("ID : %s\t Name :%s\t Grade : %d\n", user.ID, user.Name, user.Grade)

}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	respose, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer respose.Body.Close()

	err = json.NewDecoder(respose.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseUrl+"/user", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	respose, err := client.Do(request)

	if err != nil {
		return data, err
	}

	defer respose.Body.Close()

	err = json.NewDecoder(respose.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
