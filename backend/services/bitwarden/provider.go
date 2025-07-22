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
	bitwarden sdk.BitwardenClientInterface
}

func Provider() (*Client, error) {
	client, err := sdk.NewBitwardenClient(&apiUrl, &identityUrl)
	if err != nil {
		return nil, base.TagError("bitwarden.client", err)
	}

	accessToken := os.Getenv("BACKEND_BITWARDEN_TOKEN")
	if accessToken == "" {
		return &Client{bitwarden: nil}, nil
	}

	if err = client.AccessTokenLogin(accessToken, nil); err != nil {
		return nil, base.TagError("bitwarden.auth", err)
	}

	return &Client{bitwarden: client}, nil
}
