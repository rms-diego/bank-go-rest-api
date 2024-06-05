package serialize

import (
	"encoding/json"
)

func ToJson[T interface{}](dataBytes []byte) (T, error) {
	var data T

	err := json.Unmarshal(dataBytes, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
