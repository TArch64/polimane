package bitwarden

import (
	"github.com/bitwarden/sdk-go"
	"go.uber.org/fx"

	"polimane/backend/base"
	"polimane/backend/services/osenv"
	"polimane/backend/services/osfs"
)

var (
	apiUrl      = "https://api.bitwarden.eu"
	identityUrl = "https://identity.bitwarden.eu"
)

type ClientOptions struct {
	fx.In
	FS  osfs.FS
	Env osenv.Env
}

type Client interface {
	DownloadCerts(certs []*DownloadingCert) error
	Load(sids []string) (map[string]string, error)
	LoadToEnviron(names []string) error
}

type Impl struct {
	api sdk.BitwardenClientInterface
	fs  osfs.FS
	env osenv.Env
}

var _ Client = (*Impl)(nil)

func Provider(options ClientOptions) (Client, error) {
	accessToken := options.Env.Getenv("BACKEND_BITWARDEN_TOKEN")
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

	return &Impl{
		api: api,
		fs:  options.FS,
		env: options.Env,
	}, nil
}
