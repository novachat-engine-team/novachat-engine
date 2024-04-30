package util

import (
	"encoding/json"
	"novachat_engine/pkg/log"
)

func MapToJsonString(input []interface{}) string {
	output,err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(output)
}

func JsonStringToMap(inputStr string) (output []interface{},err error) {
	output = make([]interface{},1024)
	if err = json.Unmarshal([]byte(inputStr),&output); err != nil {
		return nil,err
	}
	return
}

func ConfigJsonToMap(jsonStr string,match string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Errorf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	if m,ok := m[match];ok {
		log.Infof("Load [%s] json config: %v",match,m)
		return m.(map[string]interface{}),nil
	}

	log.Warnf("Does not match [%s] from json config: %v",match,m)

	return m, nil
}