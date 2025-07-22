package bitwarden

import (
	"os"
	"path/filepath"
)

type DownloadingCert struct {
	SID  string
	Dest string
}

func (c *Client) DownloadCerts(certs []*DownloadingCert) error {
	sids := make([]string, len(certs))
	idCertMap := make(map[string]*DownloadingCert)
	for i, cert := range certs {
		sids[i] = os.Getenv(cert.SID)
		idCertMap[sids[i]] = cert
	}

	secrets, err := c.Load(sids)
	if err != nil {
		return err
	}

	var dest string
	for sid, secret := range secrets {
		dest = idCertMap[sid].Dest

		err = os.MkdirAll(filepath.Dir(dest), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(dest, []byte(secret), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
