package bitwarden

import (
	"errors"
	"testing"

	"github.com/bitwarden/sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestClient_LoadToEnviron(t *testing.T) {
	tests := []struct {
		name        string
		names       []string
		envGetenvs  map[string]string
		apiResponse *sdk.SecretsResponse
		apiError    error
		setenvError error
		wantErr     bool
	}{
		{
			name:  "successful load to environ",
			names: []string{"DB_PASSWORD", "API_KEY"},
			envGetenvs: map[string]string{
				"DB_PASSWORD_SID": "sid1",
				"API_KEY_SID":     "sid2",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "db_secret"},
					{ID: "sid2", Value: "api_secret"},
				},
			},
			apiError:    nil,
			setenvError: nil,
			wantErr:     false,
		},
		{
			name:       "empty names list",
			names:      []string{},
			envGetenvs: map[string]string{},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{},
			},
			apiError:    nil,
			setenvError: nil,
			wantErr:     false,
		},
		{
			name:  "api error during load",
			names: []string{"DB_PASSWORD"},
			envGetenvs: map[string]string{
				"DB_PASSWORD_SID": "sid1",
			},
			apiResponse: nil,
			apiError:    errors.New("api error"),
			setenvError: nil,
			wantErr:     true,
		},
		{
			name:  "setenv error",
			names: []string{"DB_PASSWORD"},
			envGetenvs: map[string]string{
				"DB_PASSWORD_SID": "sid1",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "db_secret"},
				},
			},
			apiError:    nil,
			setenvError: errors.New("setenv error"),
			wantErr:     true,
		},
		{
			name:  "single environment variable",
			names: []string{"API_KEY"},
			envGetenvs: map[string]string{
				"API_KEY_SID": "sid1",
			},
			apiResponse: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "api_secret"},
				},
			},
			apiError:    nil,
			setenvError: nil,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockEnv := &MockEnv{}
			mockSecretsManager := &MockSecretsManager{}
			mockClient := &MockBitwardenClient{}

			// Setup Getenv expectations
			for _, name := range tt.names {
				sidKey := name + "_SID"
				if sidValue, exists := tt.envGetenvs[sidKey]; exists {
					mockEnv.On("Getenv", sidKey).Return(sidValue)
				}
			}

			// Setup API expectations - always called, even for empty names list
			sids := make([]string, len(tt.names))
			for i, name := range tt.names {
				sids[i] = tt.envGetenvs[name+"_SID"]
			}
			mockClient.On("Secrets").Return(mockSecretsManager)
			mockSecretsManager.On("GetByIDS", sids).Return(tt.apiResponse, tt.apiError)

			// Setup Setenv expectations if no API error
			if tt.apiError == nil && tt.apiResponse != nil {
				for _, secret := range tt.apiResponse.Data {
					for _, name := range tt.names {
						if tt.envGetenvs[name+"_SID"] == secret.ID {
							mockEnv.On("Setenv", name, secret.Value).Return(tt.setenvError)
							break
						}
					}
				}
			}

			client := &Impl{
				api: mockClient,
				env: mockEnv,
			}

			err := client.LoadToEnviron(tt.names)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockEnv.AssertExpectations(t)
			mockClient.AssertExpectations(t)
			mockSecretsManager.AssertExpectations(t)
		})
	}
}
