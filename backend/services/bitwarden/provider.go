package bitwarden

import (
	"github.com/bitwarden/sdk-go"

	"polimane/backend/base"
	"polimane/backend/services/osenv"
	"polimane/backend/services/osfs"
)

var (
	apiUrl      = "https://api.bitwarden.eu"
	identityUrl = "https://identity.bitwarden.eu"
)

type Client struct {
	api sdk.BitwardenClientInterface
	fs  osfs.FS
	env osenv.Env
}

func Provider(fs osfs.FS, env osenv.Env) (*Client, error) {
	accessToken := env.Getenv("BACKEND_BITWARDEN_TOKEN")
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

	return &Client{
		api: api,
		fs:  fs,
		env: env,
	}, nil
}
