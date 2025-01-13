package maps

import (
	"reflect"
	"testing"
)

func TestMapsToCustomAttributes(t *testing.T) {
	type args struct {
		ignore []string
		maps   []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want CustomAttributes
	}{
		{
			name: "test1",
			args: args{
				ignore: []string{"key1"},
				maps: []map[string]interface{}{
					{"key1": "value1", "key2": "value2"},
				},
			},
			want: CustomAttributes{
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "test2",
			args: args{
				ignore: nil,
				maps: []map[string]interface{}{
					{"key1": "value1", "key2": "value2"},
				},
			},
			want: CustomAttributes{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "test3",
			args: args{
				ignore: []string{"key1"},
				maps: []map[string]interface{}{
					{"key1": "value1", "key2": []string{"value2.1", "value2.2"}},
				},
			},
			want: CustomAttributes{
				{Key: "key2", Value: "[value2.1 value2.2]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapsToCustomAttributes(tt.args.ignore, tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapsToCustomAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
