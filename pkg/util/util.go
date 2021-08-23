package util

import "encoding/json"

// map >> string
func Marshal(m map[string]interface{}) string {
	if byt, err := json.Marshal(m); err != nil {
		Errorf(err.Error())
		return ""
	} else {
		return string(byt)
	}
}

//string >> map
func Unmarshal(str string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		Errorf(err.Error())
		return nil, err
	} else {
		return data, nil
	}
}
