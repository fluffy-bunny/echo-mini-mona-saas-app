package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/labstack/echo/v4"
)

func SafeUnmarshalRequestFromBody(ctx echo.Context, v interface{}) error {
	var bodyBytes []byte
	// Read the Body content
	bodyBytes, _ = io.ReadAll(ctx.Request().Body)
	defer func() {
		ctx.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()
	err := json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}
func UnmarshalFromRequestBody(body io.Reader, v interface{}) error {
	if body == nil {
		return errors.New("body is nil")
	}

	var bodyBytes []byte
	// Read the Body content
	bodyBytes, _ = io.ReadAll(body)

	err := json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}

func StringPointer(value string) *string {
	return &value
}
