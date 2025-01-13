package maps

import "go.opentelemetry.io/collector/pdata/pcommon"

func NoopNormalisationFunc(s string) string {
	return s
}

func PCommonMapsToMap(ignore []string, normalisationFunc func(string) string, pcommonMaps ...pcommon.Map) map[string]any {
	newMap := make(map[string]any)

	for _, pcms := range pcommonMaps {
		pcms.Range(func(k string, v pcommon.Value) bool {
			newMap[normalisationFunc(k)] = v.AsString()
			return true
		})
	}

	for _, attr := range ignore {
		delete(newMap, attr)
	}

	return newMap
}

func PCommonMapsToCustomAttributes(ignore []string, pcommonMaps ...pcommon.Map) CustomAttributes {

	ig := make(map[string]struct{})
	for _, attr := range ignore {
		ig[attr] = struct{}{}
	}

	var customAttributes CustomAttributes
	for _, pcms := range pcommonMaps {
		pcms.Range(func(k string, v pcommon.Value) bool {
			if _, ok := ig[k]; ok {
				return true
			}
			customAttributes = append(customAttributes, CustomAttribute{k, v.AsString()})
			return true
		})
	}

	return customAttributes
}
