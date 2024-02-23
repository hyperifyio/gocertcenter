package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"testing"
)

func TestNewPrivateKeyDTO(t *testing.T) {
	tests := []struct {
		name           string
		serialNumber   string
		keyType        string
		privateKey     string
		wantSerial     string
		wantType       string
		wantPrivateKey string
	}{
		{
			name:           "RSA key",
			serialNumber:   "123456789",
			keyType:        "RSA",
			privateKey:     "RSA privateKey",
			wantSerial:     "123456789",
			wantType:       "RSA",
			wantPrivateKey: "RSA privateKey",
		},
		{
			name:           "ECDSA key",
			serialNumber:   "987654321",
			keyType:        "ECDSA",
			privateKey:     "ECDSA privateKey",
			wantSerial:     "987654321",
			wantType:       "ECDSA",
			wantPrivateKey: "ECDSA privateKey",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dtos.NewPrivateKeyDTO(tt.serialNumber, tt.keyType, tt.privateKey)

			if got.Certificate != tt.wantSerial {
				t.Errorf("NewPrivateKeyDTO().SerialNumber = %v, want %v", got.Certificate, tt.wantSerial)
			}

			if got.Type != tt.wantType {
				t.Errorf("NewPrivateKeyDTO().Type = %v, want %v", got.Type, tt.wantType)
			}

			if got.PrivateKey != tt.wantPrivateKey {
				t.Errorf("NewPrivateKeyDTO().PrivateKey = %v, want %v", got.PrivateKey, tt.wantPrivateKey)
			}
		})
	}
}
