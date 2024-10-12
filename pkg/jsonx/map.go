package jsonx

import (
	"encoding/json"
	"fmt"
)

func StringMap(j string) map[string]interface{} {
	mp := make(map[string]interface{})
	err := json.Unmarshal([]byte(j), &mp)
	if err != nil {
		fmt.Println(fmt.Sprintf("[jsonx.StringMap] err:%s str:%s", err.Error(), j))
	}
	return mp
}

func Value(j string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(j), &i)
	if err != nil {
		fmt.Println(fmt.Sprintf("[jsonx.Value] err:%s str:%s", err.Error(), j))
		return nil
	}
	return i
}
