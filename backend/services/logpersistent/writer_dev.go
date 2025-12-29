//go:build dev

package logpersistent

import (
	"io"
	"os"
)

func newWriter(options ProviderOptions) (io.Writer, error) {
	file, err := os.OpenFile("/tmp/app/persistent.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
