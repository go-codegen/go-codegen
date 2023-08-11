package utils

import "testing"

func TestParseCamelCaseToSnakeCase(t *testing.T) {
	type args struct {
		camelString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ParseCamelCaseToSnakeCase",
			args: args{
				camelString: "ParseCamelCaseToSnakeCase",
			},
			want: "parse_camel_case_to_snake_case",
		},
		{
			name: "already_snake_case",
			args: args{
				camelString: "already_snake_case",
			},
			want: "already_snake_case",
		},
		{
			name: "HTTPRequest",
			args: args{
				camelString: "HTTPRequest",
			},
			want: "http_request",
		},
		{
			name: "UserID",
			args: args{
				camelString: "UserID",
			},
			want: "user_id",
		},
		{
			name: "HTTP",
			args: args{
				camelString: "HTTP",
			},
			want: "http",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCamelCaseToSnakeCase(tt.args.camelString); got != tt.want {
				t.Errorf("ParseCamelCaseToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
