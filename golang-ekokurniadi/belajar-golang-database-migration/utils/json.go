package utils

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicErr(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicErr(err)
}
