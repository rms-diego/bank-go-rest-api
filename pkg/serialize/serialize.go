package serialize

import (
	"encoding/json"
	"io"
)

func BodyToJSON[T interface{}](r io.Reader) (T, error) {
	var data T

	bytesConvert, err := io.ReadAll(r)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytesConvert, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
