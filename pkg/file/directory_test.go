package file

import "testing"

func TestMakeDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeDirectory(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("MakeDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
