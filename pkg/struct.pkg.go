package pkg

import (
	"encoding/json"
	"fmt"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map

	fmt.Println(newMap)
	return
}
