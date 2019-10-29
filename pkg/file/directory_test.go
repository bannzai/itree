package file

import (
	"testing"

	"github.com/bannzai/itree/pkg/testutil"
)

func TestMakeDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfully MakeDirectory",
			args: args{
				path: testutil.TemporaryDirectoryPath(t),
			},
			wantErr: false,
		},
		{
			name: "directory is already exists",
			args: args{
				path: testutil.CurrentDirectoryPath(t) + "/testdata",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeDirectory(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("MakeDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
