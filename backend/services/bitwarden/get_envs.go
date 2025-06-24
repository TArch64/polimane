package bitwarden

import "os"

func LoadToEnviron(names []string) error {
	var id string
	var ids []string
	idNameMap := make(map[string]string)
	for _, name := range names {
		id = os.Getenv(name + "_SID")
		idNameMap[id] = name
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
