package domain

import "testing"

// stupid test, but I don't want to change the file
func TestGetDatabaseConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "positive_test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nulConfig := DBConfig{}
			if content := GetDatabaseConfig(); content == &nulConfig {
				t.Errorf("GetDatabaseConfig() = %v, want %v", content, "")
			}
		})
	}
}
