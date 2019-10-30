package parser

import (
	"reflect"
	"testing"
)

func Test_parseSize(t *testing.T) {
	type args struct {
		sizeString string
	}
	tests := []struct {
		name string
		args args
		want Size
	}{
		{
			name: "When `10 20`\n",
			args: args{
				sizeString: "10 20\n",
			},
			want: Size{
				Width:  10,
				Height: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSize(tt.args.sizeString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
