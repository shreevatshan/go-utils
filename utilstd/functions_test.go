package utilstd

import (
	"reflect"
	"testing"
)

func TestRemoveString(t *testing.T) {
	type args struct {
		str    string
		remove string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{str: "test", remove: "t"}, want: "es"},
		{name: "test", args: args{str: "test", remove: "e"}, want: "tst"},
		{name: "test", args: args{str: "test", remove: "s"}, want: "tet"},
		{name: "test", args: args{str: "this is a test", remove: "test"}, want: "this is a "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveString(&tt.args.str, tt.args.remove)
			if tt.args.str != tt.want {
				t.Errorf("RemoveString() = %v, want %v", tt.args.str, tt.want)
			}
		})
	}
}

func TestReplaceString(t *testing.T) {
	type args struct {
		str  string
		from string
		to   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{str: "test", from: "t", to: "T"}, want: "TesT"},
		{name: "test", args: args{str: "test", from: "e", to: "E"}, want: "tEst"},
		{name: "test", args: args{str: "test", from: "s", to: "S"}, want: "teSt"},
		{name: "test", args: args{str: "this is a test", from: "a test", to: "not a test"}, want: "this is not a test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReplaceString(&tt.args.str, tt.args.from, tt.args.to)
			if tt.args.str != tt.want {
				t.Errorf("ReplaceString() = %v, want %v", tt.args.str, tt.want)
			}
		})
	}
}

func TestRemoveWhiteSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{str: "This is a test"}, want: "Thisisatest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveWhiteSpace(&tt.args.str)
		})
	}
}

func TestConvertCommaSeparatedStringToSet(t *testing.T) {
	type args struct {
		comma_separated_string string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{name: "test", args: args{comma_separated_string: "1,2,3,4,5"}, want: map[string]struct{}{
			"1": struct{}{},
			"2": struct{}{},
			"3": struct{}{},
			"4": struct{}{},
			"5": struct{}{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertCommaSeparatedStringToSet(tt.args.comma_separated_string)

			for k := range tt.want {
				if got.Elements[k] != tt.want[k] {
					t.Errorf("ConvertCommaSeparatedStringToSet() = %v, want %v", got.Elements[k], tt.want[k])
				}
			}

		})
	}
}

func TestReturnKeyAndValueFromString(t *testing.T) {
	type args struct {
		keyvalue_string string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
	}{
		{name: "test", args: args{keyvalue_string: "key=value"}, wantKey: "key", wantValue: "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := ReturnKeyAndValueFromString(tt.args.keyvalue_string)
			if gotKey != tt.wantKey {
				t.Errorf("ReturnKeyAndValueFromString() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotValue != tt.wantValue {
				t.Errorf("ReturnKeyAndValueFromString() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual(t *testing.T) {
	type args struct {
		newline_separated_string string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "test", args: args{newline_separated_string: "key1=value1\nkey2=value2\nkey3=value3"}, want: map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual(tt.args.newline_separated_string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma(t *testing.T) {
	type args struct {
		newline_separated_string string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "test", args: args{newline_separated_string: "keya1,keya2,keya3=value1\nkeyb1,keyb2,keyb3=value2"}, want: map[string]string{"keya1": "value1", "keya2": "value1", "keya3": "value1", "keyb1": "value2", "keyb2": "value2", "keyb3": "value2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma(tt.args.newline_separated_string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPrefixCaseInsensitive(t *testing.T) {
	type args struct {
		string_to_check   string
		string_to_compare string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: args{string_to_check: "This is a test", string_to_compare: "this"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefixCaseInsensitive(tt.args.string_to_check, tt.args.string_to_compare); got != tt.want {
				t.Errorf("HasPrefixCaseInsensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormKeyValuePairFromMapOfValueInterface(t *testing.T) {
	type args struct {
		from_map map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []KeyValuePair
	}{
		{name: "test", args: args{from_map: map[string]interface{}{"key1": "value1", "key2": 2}}, want: []KeyValuePair{{Key: "key1", Value: "value1"}, {Key: "key2", Value: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormKeyValuePairFromMapOfValueInterface(tt.args.from_map); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormKeyValuePairFromMapOfValueInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormKeyValuePairFromMapOfValueInterfaceSlice(t *testing.T) {
	type args struct {
		from_map map[string][]interface{}
	}
	tests := []struct {
		name string
		args args
		want []KeyValuePair
	}{
		{name: "test", args: args{from_map: map[string][]interface{}{"key1": {1, 2, 3}, "key2": {"a", "b", "c"}, "key3": {"a", 2, "c"}}}, want: []KeyValuePair{{Key: "key1", Value: []interface{}{1, 2, 3}}, {Key: "key2", Value: []interface{}{"a", "b", "c"}}, {Key: "key3", Value: []interface{}{"a", 2, "c"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormKeyValuePairFromMapOfValueInterfaceSlice(tt.args.from_map); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormKeyValuePairFromMapOfValueInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueFromInterfaceMapAsString(t *testing.T) {
	type args struct {
		key           string
		interface_map map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{key: "key1", interface_map: map[string]interface{}{"key1": "value1", "key2": "value2"}}, want: "value1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueFromInterfaceMapAsString(tt.args.key, tt.args.interface_map); got != tt.want {
				t.Errorf("GetValueFromInterfaceMapAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertIntegerToBoolean(t *testing.T) {
	type args struct {
		integer int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: args{integer: 1}, want: true},
		{name: "test", args: args{integer: 0}, want: false},
		{name: "test", args: args{integer: -1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertIntegerToBoolean(tt.args.integer); got != tt.want {
				t.Errorf("ConvertIntegerToBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundTo(t *testing.T) {
	type args struct {
		n        float64
		decimals uint32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test", args: args{n: 1.23456789, decimals: 2}, want: 1.23},
		{name: "test", args: args{n: 1.23456789, decimals: 3}, want: 1.235},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundTo(tt.args.n, tt.args.decimals); got != tt.want {
				t.Errorf("RoundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ComparisonWithTolerance(t *testing.T) {
	type args struct {
		a         float64
		b         float64
		tolerance float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: args{a: 1.23456789, b: 1.23456789, tolerance: 0.00000001}, want: true},
		{name: "test", args: args{a: 1.23456789, b: 1.23456789, tolerance: 0.01}, want: true},
		{name: "test", args: args{a: 1.23456789, b: 1.23456789, tolerance: 0}, want: false},
		{name: "test", args: args{a: 0.4166666666666667, b: 0.4166666666666667, tolerance: 0.01}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ComparisonWithTolerance(tt.args.a, tt.args.b, tt.args.tolerance); got != tt.want {
				t.Errorf("Float64ComparisonWithTolerane() = %v, want %v", got, tt.want)
			}
		})
	}
}
