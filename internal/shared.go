package internal

import (
	"encoding/json"
	"errors"
	"io"
)

func UnmarshalFromRequestBody(body io.Reader, v interface{}) error {
	if body == nil {
		return errors.New("body is nil")
	}
	reuseableReader, err := ReusableReader(body)
	if err != nil {
		return err
	}

	var bodyBytes []byte
	// Read the Body content
	bodyBytes, _ = io.ReadAll(reuseableReader)

	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}

func StringPointer(value string) *string {
	return &value
}
