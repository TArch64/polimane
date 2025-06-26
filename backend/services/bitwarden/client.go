package bitwarden

import (
	"errors"
	"os"

	"github.com/bitwarden/sdk-go"

	"polimane/backend/base"
)

var (
	apiUrl      = "https://api.bitwarden.eu"
	identityUrl = "https://identity.bitwarden.eu"
)

var client sdk.BitwardenClientInterface

func Init() error {
	var err error
	client, err = sdk.NewBitwardenClient(&apiUrl, &identityUrl)
	if err != nil {
		return base.TagError("bitwarden.client", err)
	}

	accessToken := os.Getenv("BACKEND_BITWARDEN_TOKEN")
	if accessToken == "" {
		return errors.New("'BACKEND_BITWARDEN_TOKEN' environment variable must be set")
	}

	err = client.AccessTokenLogin(accessToken, nil)
	return base.TagError("bitwarden.auth", err)
}
