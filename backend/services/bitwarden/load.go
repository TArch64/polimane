package bitwarden

func (c *Client) Load(sids []string) (map[string]string, error) {
	res, err := c.bitwarden.Secrets().GetByIDS(sids)
	if err != nil {
		return nil, err
	}

	secrets := make(map[string]string, len(res.Data))
	for _, secret := range res.Data {
		secrets[secret.ID] = secret.Value
	}

	return secrets, nil
}
