package crypt

import (
	"testing"
)

func TestDecryptAES(t *testing.T) {
	type args struct {
		cipherText string
		secretKey  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test EncryptAES",
			args: args{
				cipherText: "UwSPAIHPncPdFZ91fbC5qtYX/ReRcJzXpwQkTs+whiY=",
				secretKey:  "scg543dhjkfghjkl",
			},
			want: "This1Is2a3Test_Plain5Text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptAES(tt.args.cipherText, tt.args.secretKey)
			if err != nil {
				t.Errorf("DecryptAES() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptAES(t *testing.T) {
	type args struct {
		plainText string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test EncryptAES",
			args: args{
				plainText: "This1Is2a3Test_Plain5Text",
				secretKey: "scg543dhjkfghjkl",
			},
			want: "UwSPAIHPncPdFZ91fbC5qtYX/ReRcJzXpwQkTs+whiY=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptAES(tt.args.plainText, tt.args.secretKey)
			if err != nil {
				t.Errorf("EncryptAES() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}
