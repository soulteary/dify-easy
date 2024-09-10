package Fn

import (
	"bytes"
	"encoding/json"
)

func ConvertArrToCommand(arr []string) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(arr)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
