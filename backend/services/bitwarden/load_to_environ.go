package bitwarden

import "os"

func (c *Client) LoadToEnviron(names []string) error {
	ids := make([]string, len(names))
	idNameMap := make(map[string]string)
	for i, name := range names {
		ids[i] = os.Getenv(name + "_SID")
		idNameMap[ids[i]] = name
	}

	secrets, err := c.Load(ids)
	if err != nil {
		return err
	}

	for sid, secret := range secrets {
		if err = os.Setenv(idNameMap[sid], secret); err != nil {
			return err
		}
	}

	return nil
}
