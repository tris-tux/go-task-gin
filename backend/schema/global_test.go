package schema_test

import (
	"reflect"
	"testing"

	. "github.com/tris-tux/go-task-gin/backend/schema"
)

func TestErrorWrap(t *testing.T) {
	type args struct {
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
		want *WrappedError
	}{
		{
			name: "aa",
			args: args{
				code:    404,
				message: "Data Not Found",
			},
			want: &WrappedError{
				Code:    404,
				Message: "Data Not Found",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorWrap(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorWrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
