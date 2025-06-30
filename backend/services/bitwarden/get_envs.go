package bitwarden

import "os"

func LoadToEnviron(names []string) error {
	ids := make([]string, len(names))
	idNameMap := make(map[string]string)
	for i, name := range names {
		ids[i] = os.Getenv(name + "_SID")
		idNameMap[ids[i]] = name
	}

	res, err := client.Secrets().GetByIDS(ids)
	if err != nil {
		return err
	}

	for _, secret := range res.Data {
		err = os.Setenv(idNameMap[secret.ID], secret.Value)
		if err != nil {
			return err
		}
	}

	return nil
}
