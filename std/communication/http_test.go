package communication

import (
	"os"
	"reflect"
	"testing"
)

func TestGzip(t *testing.T) {
	type args struct {
		inputfile  string
		outputfile string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test1",
			args: args{
				inputfile:  "testdata/testfile",
				outputfile: "testdata/testfile.gz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b, _ := os.ReadFile(tt.args.inputfile)

			got, err := Gzip(b)

			//os.WriteFile(tt.args.outputfile, got, 0644)

			tt.want, _ = os.ReadFile(tt.args.outputfile)

			if (err != nil) != tt.wantErr {
				t.Errorf("Gzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gzip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractGzip(t *testing.T) {
	type args struct {
		inputfile  string
		outputfile string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test1",
			args: args{
				inputfile:  "testdata/testfile.gz",
				outputfile: "testdata/testfile",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b, _ := os.ReadFile(tt.args.inputfile)

			tt.want, _ = os.ReadFile(tt.args.outputfile)

			got, err := ExtractGzip(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractGzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractGzip() = %v, want %v", got, tt.want)
			}
		})
	}
}
