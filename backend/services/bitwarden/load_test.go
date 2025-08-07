package bitwarden

import (
	"errors"
	"testing"

	"github.com/bitwarden/sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestClient_Load(t *testing.T) {
	tests := []struct {
		name     string
		sids     []string
		response *sdk.SecretsResponse
		apiError error
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "successful load with multiple secrets",
			sids: []string{"sid1", "sid2"},
			response: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "secret1"},
					{ID: "sid2", Value: "secret2"},
				},
			},
			apiError: nil,
			expected: map[string]string{
				"sid1": "secret1",
				"sid2": "secret2",
			},
			wantErr: false,
		},
		{
			name:     "empty sids list",
			sids:     []string{},
			response: &sdk.SecretsResponse{Data: []sdk.SecretResponse{}},
			apiError: nil,
			expected: map[string]string{},
			wantErr:  false,
		},
		{
			name:     "api error",
			sids:     []string{"sid1"},
			response: nil,
			apiError: errors.New("api error"),
			expected: nil,
			wantErr:  true,
		},
		{
			name: "single secret",
			sids: []string{"sid1"},
			response: &sdk.SecretsResponse{
				Data: []sdk.SecretResponse{
					{ID: "sid1", Value: "secret1"},
				},
			},
			apiError: nil,
			expected: map[string]string{
				"sid1": "secret1",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSecretsManager := &MockSecretsManager{}
			mockClient := &MockBitwardenClient{}

			mockClient.On("Secrets").Return(mockSecretsManager)
			mockSecretsManager.On("GetByIDS", tt.sids).Return(tt.response, tt.apiError)

			client := &Impl{
				api: mockClient,
			}

			result, err := client.Load(tt.sids)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			mockClient.AssertExpectations(t)
			mockSecretsManager.AssertExpectations(t)
		})
	}
}
