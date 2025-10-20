package bitwarden

import (
	"os"

	"github.com/bitwarden/sdk-go"
	"go.uber.org/fx"

	"polimane/backend/base"
)

var (
	apiUrl      = "https://api.bitwarden.eu"
	identityUrl = "https://identity.bitwarden.eu"
)

type ClientOptions struct {
	fx.In
}

type Client interface {
	DownloadCerts(certs []*DownloadingCert) error
	Load(sids []string) (map[string]string, error)
	LoadToEnviron(names []string) error
}

type Impl struct {
	api sdk.BitwardenClientInterface
}

var _ Client = (*Impl)(nil)

func Provider(options ClientOptions) (Client, error) {
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

	return &Impl{api: api}, nil
}
