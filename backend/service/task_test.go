package service_test

import (
	"testing"

	"github.com/tris-tux/go-task-gin/backend/schema"
	. "github.com/tris-tux/go-task-gin/backend/service"
)

func Test_isFinishedTask(t *testing.T) {
	type args struct {
		detail []schema.DetailResponse
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "false",
			args: args{
				detail: []schema.DetailResponse{
					{
						ObjectName: "asd",
						IsFinished: false,
					},
					{
						ObjectName: "fds",
						IsFinished: false,
					},
					{
						ObjectName: "gdf",
						IsFinished: false,
					},
				},
			},
			want: false,
		},
		{
			name: "falsetrue",
			args: args{
				detail: []schema.DetailResponse{
					{
						ObjectName: "asd",
						IsFinished: false,
					},
					{
						ObjectName: "fds",
						IsFinished: true,
					},
					{
						ObjectName: "gdf",
						IsFinished: true,
					},
				},
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				detail: []schema.DetailResponse{
					{
						ObjectName: "asd",
						IsFinished: true,
					},
					{
						ObjectName: "fds",
						IsFinished: true,
					},
					{
						ObjectName: "gdf",
						IsFinished: true,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFinishedTask(tt.args.detail); got != tt.want {
				t.Errorf("isFinishedTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
