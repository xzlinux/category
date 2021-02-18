package common

import (
	"encoding/json"
	"fmt"
)

func SwapTo(request, category interface{}) (err error) {
	fmt.Println(request)
	dateByte, err := json.Marshal(request)

	if err != nil {
		return
	}
	err = json.Unmarshal(dateByte, category)
	fmt.Println("datrt")
	fmt.Println(category)
	fmt.Println("end")
	return err
}
