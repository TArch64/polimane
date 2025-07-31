package bitwarden

import (
	"errors"
	"os"
	"testing"

	"github.com/bitwarden/sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestClient_DownloadCerts(t *testing.T) {
	tests := []struct {
		name        string
		certs       []*DownloadingCert
		envGetenvs  map[string]string
		apiResponse *sdk.SecretsResponse
		apiError    error
		mkdirError  error
		writeError  error
		wantErr     bool
	}{
		{
			name: "successful download of multiple certs",
			certs: []*DownloadingCert{
				{SID: "CERT1_SID", Dest: "/tmp/cert1.pem"},
				{SID: "CERT2_SID", Dest: "/tmp/cert2.pem"},
			},
			envGetenvs: map[string]string{
				"CERT1_SID": "sid1",
				"CERT2_SID": "sid2",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "cert1-content"},
					{ID: "sid2", Value: "cert2-content"},
				},
			},
			apiError:   nil,
			mkdirError: nil,
			writeError: nil,
			wantErr:    false,
		},
		{
			name:       "empty certs list",
			certs:      []*DownloadingCert{},
			envGetenvs: map[string]string{},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{},
			},
			apiError:   nil,
			mkdirError: nil,
			writeError: nil,
			wantErr:    false,
		},
		{
			name: "api error during load",
			certs: []*DownloadingCert{
				{SID: "CERT1_SID", Dest: "/tmp/cert1.pem"},
			},
			envGetenvs: map[string]string{
				"CERT1_SID": "sid1",
			},
			apiResponse: nil,
			apiError:    errors.New("api error"),
			mkdirError:  nil,
			writeError:  nil,
			wantErr:     true,
		},
		{
			name: "mkdir error",
			certs: []*DownloadingCert{
				{SID: "CERT1_SID", Dest: "/tmp/cert1.pem"},
			},
			envGetenvs: map[string]string{
				"CERT1_SID": "sid1",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "cert1-content"},
				},
			},
			apiError:   nil,
			mkdirError: errors.New("mkdir error"),
			writeError: nil,
			wantErr:    true,
		},
		{
			name: "write file error",
			certs: []*DownloadingCert{
				{SID: "CERT1_SID", Dest: "/tmp/cert1.pem"},
			},
			envGetenvs: map[string]string{
				"CERT1_SID": "sid1",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "cert1-content"},
				},
			},
			apiError:   nil,
			mkdirError: nil,
			writeError: errors.New("write error"),
			wantErr:    true,
		},
		{
			name: "single cert with nested directory",
			certs: []*DownloadingCert{
				{SID: "CERT1_SID", Dest: "/tmp/certs/ssl/cert.pem"},
			},
			envGetenvs: map[string]string{
				"CERT1_SID": "sid1",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "cert-content"},
				},
			},
			apiError:   nil,
			mkdirError: nil,
			writeError: nil,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockEnv := &MockEnv{}
			mockFS := &MockFS{}
			mockSecretsManager := &MockSecretsManager{}
			mockClient := &MockBitwardenClient{}

			// Setup Getenv expectations
			for _, cert := range tt.certs {
				if sidValue, exists := tt.envGetenvs[cert.SID]; exists {
					mockEnv.On("Getenv", cert.SID).Return(sidValue)
				}
			}

			// Setup API expectations - always called, even for empty cert lists
			sids := make([]string, len(tt.certs))
			for i, cert := range tt.certs {
				sids[i] = tt.envGetenvs[cert.SID]
			}
			mockClient.On("Secrets").Return(mockSecretsManager)
			mockSecretsManager.On("GetByIDS", sids).Return(tt.apiResponse, tt.apiError)

			// Setup filesystem expectations if no API error
			if tt.apiError == nil && tt.apiResponse != nil {
				for _, secret := range tt.apiResponse.Data {
					for _, cert := range tt.certs {
						if tt.envGetenvs[cert.SID] == secret.ID {
							mockFS.On("MkdirAll", "/tmp", os.ModePerm).Return(tt.mkdirError).Maybe()
							mockFS.On("MkdirAll", "/tmp/certs/ssl", os.ModePerm).Return(tt.mkdirError).Maybe()
							if tt.mkdirError == nil {
								mockFS.On("WriteFile", cert.Dest, []byte(secret.Value), os.FileMode(0644)).Return(tt.writeError)
							}
							break
						}
					}
				}
			}

			client := &Client{
				api: mockClient,
				fs:  mockFS,
				env: mockEnv,
			}

			err := client.DownloadCerts(tt.certs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockEnv.AssertExpectations(t)
			mockFS.AssertExpectations(t)
			mockClient.AssertExpectations(t)
			mockSecretsManager.AssertExpectations(t)
		})
	}
}
