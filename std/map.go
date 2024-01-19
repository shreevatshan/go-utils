package std

import (
	"crypto/sha256"
	"encoding/json"
)

type KeyValuePair struct {
	Key   string
	Value interface{}
}

func FormKeyValuePairFromMapOfValueInterface(fromMap map[string]interface{}) []KeyValuePair {
	var toArray []KeyValuePair
	for key, value := range fromMap {
		var pair KeyValuePair
		pair.Key = key
		pair.Value = value
		toArray = append(toArray, pair)
	}
	return toArray
}

func FormKeyValuePairFromMapOfValueInterfaceSlice(fromMap map[string][]interface{}) []KeyValuePair {
	var toArray []KeyValuePair
	for key, value := range fromMap {
		var pair KeyValuePair
		pair.Key = key
		pair.Value = value
		toArray = append(toArray, pair)
	}
	return toArray
}

func GetValueFromInterfaceMapAsString(key string, interfaceMap map[string]interface{}) string {
	var value string
	if _, exists := interfaceMap[key]; exists {
		value = interfaceMap[key].(string)
	}
	return value
}

func MapToHashCode(value interface{}) ([]byte, error) {

	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	return hash[:], nil
}

// converts the given struct to map[string]interface{}, where the key is the json field name
func StructToMap(any interface{}) (map[string]interface{}, error) {
	newMap := make(map[string]interface{})
	data, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return nil, err
	}
	return newMap, nil
}
