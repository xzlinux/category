package common

import "encoding/json"

func SwapTo(request, category interface{}) (err error) {
	dateByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dateByte, category)
}
