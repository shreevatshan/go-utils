package std

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
		commaSeparatedString string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{name: "test", args: args{commaSeparatedString: "1,2,3,4,5"}, want: map[string]struct{}{
			"1": struct{}{},
			"2": struct{}{},
			"3": struct{}{},
			"4": struct{}{},
			"5": struct{}{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertCommaSeparatedStringToSet(tt.args.commaSeparatedString)

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
		keyvalueString string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
	}{
		{name: "test", args: args{keyvalueString: "key=value"}, wantKey: "key", wantValue: "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := ReturnKeyAndValueFromString(tt.args.keyvalueString)
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
		newlineSeparatedString string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "test", args: args{newlineSeparatedString: "key1=value1\nkey2=value2\nkey3=value3"}, want: map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual(tt.args.newlineSeparatedString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma(t *testing.T) {
	type args struct {
		newlineSeparatedString string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "test", args: args{newlineSeparatedString: "keya1,keya2,keya3=value1\nkeyb1,keyb2,keyb3=value2"}, want: map[string]string{"keya1": "value1", "keya2": "value1", "keya3": "value1", "keyb1": "value2", "keyb2": "value2", "keyb3": "value2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma(tt.args.newlineSeparatedString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPrefixCaseInsensitive(t *testing.T) {
	type args struct {
		stringToCheck   string
		stringToCompare string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: args{stringToCheck: "This is a test", stringToCompare: "this"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefixCaseInsensitive(tt.args.stringToCheck, tt.args.stringToCompare); got != tt.want {
				t.Errorf("HasPrefixCaseInsensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
