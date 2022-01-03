package utils

import "encoding/json"

func InterfaceToDataStruct(i interface{}, data interface{}) error {
	iJson, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return json.Unmarshal(iJson, data)
}
