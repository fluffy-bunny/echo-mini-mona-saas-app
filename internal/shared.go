package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SafeGetBodyFromHttpResponse(response *http.Response) ([]byte, error) {
	var bodyBytes []byte
	// Read the Body content
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()
	return bodyBytes, nil
}
func SafeGetBodyFromHttpRequest(request *http.Request) ([]byte, error) {
	var bodyBytes []byte
	// Read the Body content
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyBytes))
	defer func() {
		request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()
	return bodyBytes, nil
}
func SafeUnmarshalFromHttpResponse(response *http.Response, v interface{}) error {
	var bodyBytes []byte
	// Read the Body content
	bodyBytes, err := SafeGetBodyFromHttpResponse(response)
	if err != nil {
		return err
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		return err
	}
	return nil
}
func SafeUnmarshalFromHttpRequest(request *http.Request, v interface{}) error {
	var bodyBytes []byte
	// Read the Body content
	bodyBytes, err := SafeGetBodyFromHttpRequest(request)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}

func StringPointer(value string) *string {
	return &value
}
