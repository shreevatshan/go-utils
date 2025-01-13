package file

import (
	"reflect"
	"testing"
)

func TestListByModTime(t *testing.T) {
	type args struct {
		dirname string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test1",
			args: args{
				dirname: "testdata",
			},
			want: []string{
				"testdata/testfile1",
				"testdata/testfile2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListByModTime(tt.args.dirname)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByModTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListByModTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
