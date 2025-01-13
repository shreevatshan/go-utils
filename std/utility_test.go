package std

import (
	"reflect"
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

func TestDecodeBase64(t *testing.T) {
	type args struct {
		encoded []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{encoded: []byte("dXNfMDFmY2I4MzQzYmE4NDJiMmU0NzZiZDM1NDc5NzJmZmU=")},
			want:    []byte("us_01fcb8343ba842b2e476bd3547972ffe"),
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{encoded: []byte("")},
			want:    []byte(""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.args.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
