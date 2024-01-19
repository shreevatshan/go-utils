package std

import (
	"reflect"
	"testing"
)

func TestFormKeyValuePairFromMapOfValueInterface(t *testing.T) {
	type args struct {
		fromMap map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []KeyValuePair
	}{
		{name: "test", args: args{fromMap: map[string]interface{}{"key1": "value1", "key2": 2}}, want: []KeyValuePair{{Key: "key1", Value: "value1"}, {Key: "key2", Value: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormKeyValuePairFromMapOfValueInterface(tt.args.fromMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormKeyValuePairFromMapOfValueInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormKeyValuePairFromMapOfValueInterfaceSlice(t *testing.T) {
	type args struct {
		fromMap map[string][]interface{}
	}
	tests := []struct {
		name string
		args args
		want []KeyValuePair
	}{
		{name: "test", args: args{fromMap: map[string][]interface{}{"key1": {1, 2, 3}, "key2": {"a", "b", "c"}, "key3": {"a", 2, "c"}}}, want: []KeyValuePair{{Key: "key1", Value: []interface{}{1, 2, 3}}, {Key: "key2", Value: []interface{}{"a", "b", "c"}}, {Key: "key3", Value: []interface{}{"a", 2, "c"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormKeyValuePairFromMapOfValueInterfaceSlice(tt.args.fromMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormKeyValuePairFromMapOfValueInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueFromInterfaceMapAsString(t *testing.T) {
	type args struct {
		key          string
		interfaceMap map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{key: "key1", interfaceMap: map[string]interface{}{"key1": "value1", "key2": "value2"}}, want: "value1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueFromInterfaceMapAsString(tt.args.key, tt.args.interfaceMap); got != tt.want {
				t.Errorf("GetValueFromInterfaceMapAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
