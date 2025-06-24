package bitwarden

import (
	"errors"
	"os"

	"github.com/bitwarden/sdk-go"
)

var (
	apiUrl      = "https://api.bitwarden.com"
	identityUrl = "https://identity.bitwarden.com"
)

var client sdk.BitwardenClientInterface

func Init() error {
	var err error
	client, err = sdk.NewBitwardenClient(&apiUrl, &identityUrl)
	if err != nil {
		return err
	}

	accessToken := os.Getenv("BACKEND_BITWARDEN_TOKEN")
	if accessToken == "" {
		return errors.New("'BACKEND_BITWARDEN_TOKEN' environment variable must be set")
	}

	return client.AccessTokenLogin(accessToken, nil)
}
