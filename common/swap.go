package common

import (
	"encoding/json"
	"fmt"
)

func SwapTo(request, category interface{}) (err error) {
	fmt.Printf("request type is:", request)
	fmt.Printf("request type is: %T\n", request)
	dateByte, err := json.Marshal(request)
	fmt.Printf("request:%s", dateByte)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dateByte, category)
	fmt.Println("datrt")
	fmt.Println(category)
	fmt.Println("end")
	dateByte, err = json.Marshal(category)
	fmt.Printf("request222:", string(dateByte))
	return err
}
