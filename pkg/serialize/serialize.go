package serialize

import (
	"encoding/json"
	"io"
)

func BodyToJSON[T interface{}](dataReader io.Reader) (T, error) {
	var data T

	bytesConvert, err := io.ReadAll(dataReader)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytesConvert, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
