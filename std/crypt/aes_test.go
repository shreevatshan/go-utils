package crypt

import (
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	type args struct {
		plainText []byte
		secretKey []byte
		iv        []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test1 EncryptAES",
			args: args{
				plainText: []byte("us_01fcb8343ba842b2e476bd3547972ffe"),
				secretKey: []byte("20240919T16050691154561500000000"),
				iv:        []byte("8607468918117109"),
			},
			want: "JVA21CoWjfUI4wXhzfom75WmI8Erzjq78dzkCHPo3tx1r2flHtoOiQW3wZSq87nw",
		},
		{name: "Test2 EncryptAES",
			args: args{
				plainText: []byte("us_01fcb8343ba842b2e476bd3547972ffe"),
				secretKey: []byte("20240919T16050691154561500000000"),
				iv:        []byte("860746891811710"),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESEncrypt(tt.args.plainText, tt.args.secretKey, tt.args.iv)
			if err != nil && !tt.wantErr {
				t.Errorf("EncryptAES() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESDecrypt(t *testing.T) {
	type args struct {
		cipherText []byte
		secretKey  []byte
		iv         []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test1 EncryptAES",
			args: args{
				cipherText: []byte("JVA21CoWjfUI4wXhzfom75WmI8Erzjq78dzkCHPo3tx1r2flHtoOiQW3wZSq87nw"),
				secretKey:  []byte("20240919T16050691154561500000000"),
				iv:         []byte("8607468918117109"),
			},
			want: "us_01fcb8343ba842b2e476bd3547972ffe",
		},
		{
			name: "Test1 EncryptAES",
			args: args{
				cipherText: []byte("eefc4fabe0c2b78ecfdc8cdcbe676724"),
				secretKey:  []byte("20240919T16050691154561500000000"),
				iv:         []byte("8607468918117109"),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESDecrypt(tt.args.cipherText, tt.args.secretKey, tt.args.iv)
			if err != nil && !tt.wantErr {
				t.Errorf("EncryptAES() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}
