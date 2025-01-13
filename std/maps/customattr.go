package maps

type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CustomAttributes []CustomAttribute

func (ca CustomAttributes) ToMap() map[string]string {
	m := make(map[string]string)
	for _, attr := range ca {
		m[attr.Key] = attr.Value
	}
	return m
}
