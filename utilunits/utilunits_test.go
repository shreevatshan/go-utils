package utilunits

import (
	"testing"
)

func TestConvertToNumeric(t *testing.T) {
	type args struct {
		value float64
		unit  string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "test", args: args{value: 1, unit: "hundred"}, want: 100, wantErr: false},
		{name: "test", args: args{value: 1, unit: "thousand"}, want: 1000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "lakh"}, want: 100000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "million"}, want: 1000000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "billion"}, want: 1000000000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "trillion"}, want: 1000000000000, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToNumeric(tt.args.value, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToNumeric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToBytes(t *testing.T) {
	type args struct {
		value float64
		unit  string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "test", args: args{value: 1, unit: "bytes"}, want: 1, wantErr: false},
		{name: "test", args: args{value: 1, unit: "B"}, want: 1, wantErr: false},
		{name: "test", args: args{value: 1, unit: "KB"}, want: 1024, wantErr: false},
		{name: "test", args: args{value: 1, unit: "mb"}, want: 1048576, wantErr: false},
		{name: "test", args: args{value: 1, unit: "Gb"}, want: 1073741824, wantErr: false},
		{name: "test", args: args{value: 1, unit: "tB"}, want: 1099511627776, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToBytes(tt.args.value, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToMilliseconds(t *testing.T) {
	type args struct {
		value float64
		unit  string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "test", args: args{value: 1, unit: "µs"}, want: 0.001, wantErr: false},
		{name: "test", args: args{value: 1, unit: "ms"}, want: 1, wantErr: false},
		{name: "test", args: args{value: 1, unit: "sec"}, want: 1000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "min"}, want: 60000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "hr"}, want: 3600000, wantErr: false},
		{name: "test", args: args{value: 1, unit: "day"}, want: 86400000, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToMilliseconds(tt.args.value, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToMilliseconds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToMilliseconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
