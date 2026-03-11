package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type D map[string]interface{}

func main() {

	baseUrl := "http://localhost:9000"
	method := "POST"
	data := D{"name": "bay bay"}
	responseBody, err := doRequest(baseUrl+"/data", method, data)
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}

	log.Printf("%#v \n", responseBody)

}

func doRequest(url, method string, data interface{}) (interface{}, error) {
	var payload *bytes.Buffer = nil

	if data != nil {
		payload = new(bytes.Buffer)
		err := json.NewEncoder(payload).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	certFile, err := ioutil.ReadFile("server.crt")
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)

	tlsConfig := &tls.Config{RootCAs: caCertPool}
	tlsConfig.BuildNameToCertificate()

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	responseBody := make(D)
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
