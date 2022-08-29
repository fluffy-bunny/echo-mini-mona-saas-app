package internal

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

func UnmarshalFromRequestBody(ctx echo.Context, v interface{}) error {
	reuseableReader, err := ReusableReader(ctx.Request().Body)
	if err != nil {
		return err
	}

	var bodyBytes []byte
	// Read the Body content

	if ctx.Request().Body != nil {
		bodyBytes, _ = io.ReadAll(reuseableReader)
	}

	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}
