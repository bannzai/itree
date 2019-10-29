package file

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testdata/testfile.go is exists",
			args: args{
				filename: currentDirectoryPath(t) + "/testdata/testfile.go",
			},
			want: true,
		},
		{
			name: "testdata/_testfile.go is not exists",
			args: args{
				filename: currentDirectoryPath(t) + "/testdata/_testfile.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveFile(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfully move file",
			args: args{
				from: pathForCreatedTemporaryFile(t),
				to:   temporaryFilePath(t),
			},
			wantErr: false,
		},
		{
			name: "when to path is already exists",
			args: args{
				from: pathForCreatedTemporaryFile(t),
				to:   currentDirectoryPath(t) + "/testdata/testfile.go",
			},
			wantErr: true,
		},
		{
			name: "when from path is not exists",
			args: args{
				from: currentDirectoryPath(t) + "/testdata/_testfile.go",
				to:   temporaryFilePath(t),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MoveFile(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("MoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMakeFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfully make file",
			args: args{
				path: temporaryFilePath(t),
			},
			wantErr: false,
		},
		{
			name: "file is already exists",
			args: args{
				path: currentDirectoryPath(t) + "/testdata/testfile.go",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("MakeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
