package domain

import (
	"testing"
)

func TestValidator_ValidateCreate(t *testing.T) {
	tests := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &Request{
				Name: "test",
				Age:  20,
			},
			wantErr: false,
		},
		{
			name: "negative_test_name",
			req: &Request{
				Name: "",
				Age:  20,
			},
			wantErr: true,
		},
		{
			name: "negative_test_age",
			req: &Request{
				Name: "test",
				Age:  -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			if err := v.ValidateCreate(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestValidator_ValidateChangeAge(t *testing.T) {
	tests := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &Request{
				Age: 20,
			},
			wantErr: false,
		},
		{
			name: "negative_test",
			req: &Request{
				Age: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			if err := v.ValidateChangeAge(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateChangeAge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_ValidateMakeFriends(t *testing.T) {
	tests := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &Request{
				SourceID: 1,
				TargetID: 2,
			},
			wantErr: false,
		},
		{
			name: "negative_test_same_input",
			req: &Request{
				SourceID: 1,
				TargetID: 1,
			},
			wantErr: true,
		},
		{
			name: "negative_test_negative_input",
			req: &Request{
				SourceID: -1,
				TargetID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			if err := v.ValidateMakeFriends(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateMakeFriends() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_ValidateDeleteUser(t *testing.T) {
	test := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &Request{
				TargetID: 1,
			},
			wantErr: false,
		},
		{
			name: "negative_test",
			req: &Request{
				TargetID: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			if err := v.ValidateDeleteUser(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateDeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_ValidateGetFriends(t *testing.T) {
	test := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &Request{
				TargetID: 1,
			},
			wantErr: false,
		},
		{
			name: "negative_test",
			req: &Request{
				TargetID: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			if err := v.ValidateGetFriends(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateGetFriends() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
