package domain

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRequest_Bind(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive_test",
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(bytes.NewBufferString(`{"name":"test","age":20,"friends":[]}`)),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req Request
			if err := req.Bind(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Request.Bind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
