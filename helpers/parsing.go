package helpers

import "encoding/json"

// MapToStruct method
func MapToStruct(data map[string]interface{}, target interface{}) error {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &target)
	if err != nil {
		return err
	}

	return nil
}
