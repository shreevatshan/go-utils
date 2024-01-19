package std

import (
	"testing"
)

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
