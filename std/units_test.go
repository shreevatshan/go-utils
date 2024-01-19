package std

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
