package handler_test

import (
	"errors"
	"testing"

	. "github.com/tris-tux/go-task-gin/backend/handler"
)

func Test_errorCode(t *testing.T) {
	type args struct {
		er error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "202 string",
			args: args{
				er: errors.New("202 asdaf"),
			},
			want: 202,
		},
		{
			name: "202 string",
			args: args{
				er: errors.New("404 asdaf"),
			},
			want: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorCode(tt.args.er); got != tt.want {
				t.Errorf("ErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
