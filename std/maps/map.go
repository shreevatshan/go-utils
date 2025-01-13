package maps

import (
	"encoding/json"
	"fmt"
)

// converts the given struct to map[string]interface{}, where the key is the json field name
func StructToMap(any interface{}) (map[string]interface{}, error) {
	var err error
	data, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}

	newMap := make(map[string]interface{})
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return nil, err
	}
	return newMap, nil
}

func MapsToCustomAttributes(ignore []string, maps ...map[string]interface{}) CustomAttributes {

	ig := make(map[string]struct{})
	for _, attr := range ignore {
		ig[attr] = struct{}{}
	}

	var customAttributes CustomAttributes
	for _, m := range maps {
		for k, v := range m {
			if _, ok := ig[k]; !ok {
				customAttributes = append(customAttributes, CustomAttribute{
					Key:   k,
					Value: fmt.Sprintf("%v", v),
				})
			}
		}
	}
	return customAttributes
}
