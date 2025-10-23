package bitwarden

import (
	"os"

	"github.com/bitwarden/sdk-go"

	"polimane/backend/base"
)

var (
	apiUrl      = "https://api.bitwarden.eu"
	identityUrl = "https://identity.bitwarden.eu"
)

type Client struct {
	api sdk.BitwardenClientInterface
}

func Provider() (*Client, error) {
	accessToken := os.Getenv("BACKEND_BITWARDEN_TOKEN")
	if accessToken == "" {
		return nil, nil
	}

	api, err := sdk.NewBitwardenClient(&apiUrl, &identityUrl)
	if err != nil {
		return nil, base.TagError("bitwarden.client", err)
	}

	if err = api.AccessTokenLogin(accessToken, nil); err != nil {
		return nil, base.TagError("bitwarden.auth", err)
	}

	return &Client{api: api}, nil
}
