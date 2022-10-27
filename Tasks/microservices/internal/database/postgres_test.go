package database

import (
	"log"
	"testing"
	"user/internal/domain"
	u "user/internal/user"
)

// strange to test this level, because it's called when everything is tested,
// and in fact it's just puts the data into the database
func TestClient_CreateDeleteUser(t *testing.T) {
	test := []struct {
		name    string
		model   u.User
		wantErr bool
	}{
		{
			name: "positive_test",
			model: u.User{
				Name: "test",
				Age:  20,
			},
			wantErr: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			databaseConfig := domain.GetDatabaseConfig()
			client, err := NewClient(databaseConfig, log.Default())
			if err != nil {
				t.Errorf("Client.NewClient() error = %v", err)
			}
			id, err := client.CreateUser(tt.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := client.DeleteUser(id); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateUserChangeAge(t *testing.T) {
	test := []struct {
		name    string
		model   u.User
		newAge  int
		wantErr bool
	}{
		{
			name: "positive_test",
			model: u.User{
				Name: "test",
				Age:  20,
			},
			newAge:  30,
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			databaseConfig := domain.GetDatabaseConfig()
			client, err := NewClient(databaseConfig, log.Default())
			if err != nil {
				t.Errorf("Client.NewClient() error = %v", err)
			}
			id, err := client.CreateUser(tt.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := client.ChangeAge(id, tt.newAge); (err != nil) != tt.wantErr {
				t.Errorf("Client.ChangeAge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := client.DeleteUser(id); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetUsers(t *testing.T) {
	test := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "positive_test",
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			databaseConfig := domain.GetDatabaseConfig()
			client, err := NewClient(databaseConfig, log.Default())
			if err != nil {
				t.Errorf("Client.NewClient() error = %v", err)
			}
			if _, err := client.GetUsers(); (err != nil) != tt.wantErr {
				t.Errorf("Client.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetUser(t *testing.T) {
	test := []struct {
		name    string
		user    u.User
		wantErr bool
	}{
		{
			name: "positive_test",
			user: u.User{
				Name: "test",
				Age:  20,
			},
			wantErr: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			databaseConfig := domain.GetDatabaseConfig()
			client, err := NewClient(databaseConfig, log.Default())
			if err != nil {
				t.Errorf("Client.NewClient() error = %v", err)
			}
			id, err := client.CreateUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := client.GetUserByID(id); (err != nil) != tt.wantErr {
				t.Errorf("Client.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := client.DeleteUser(id); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_MakeFriends(t *testing.T) {
	tests := []struct {
		name    string
		id1     int
		id2     int
		wantErr bool
	}{
		{
			name:    "positive_test",
			id1:     1,
			id2:     2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseConfig := domain.GetDatabaseConfig()
			client, err := NewClient(databaseConfig, log.Default())
			if err != nil {
				t.Errorf("Client.NewClient() error = %v", err)
			}
			if err := client.MakeFriends(tt.id1, tt.id2); (err != nil) != tt.wantErr {
				t.Errorf("Client.MakeFriends() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
