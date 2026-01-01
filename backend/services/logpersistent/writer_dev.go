//go:build dev

package logpersistent

import (
	"io"
	"os"
)

func newWriter(_ ProviderOptions) (io.Writer, error) {
	return os.Stdout, nil
}
