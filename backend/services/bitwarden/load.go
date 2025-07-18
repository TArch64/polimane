package bitwarden

func Load(sids []string) (map[string]string, error) {
	res, err := client.Secrets().GetByIDS(sids)
	if err != nil {
		return nil, err
	}

	secrets := make(map[string]string, len(res.Data))
	for _, secret := range res.Data {
		secrets[secret.ID] = secret.Value
	}

	return secrets, nil
}
