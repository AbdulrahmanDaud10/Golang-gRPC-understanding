package configs

import "testing"

func TestEnvMongoURI(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnvMongoURI(); got != tt.want {
				t.Errorf("EnvMongoURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
