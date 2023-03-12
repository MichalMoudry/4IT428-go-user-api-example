package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func UnmarshalRequest(request *http.Request, b any) error {
	body, error := io.ReadAll(request.Body)
	if error != nil {
		return error
	}

	if error := json.Unmarshal(body, b); error != nil {
		return error
	}

	if validate.Struct(b); error != nil {
		return error
	}

	return nil
}
