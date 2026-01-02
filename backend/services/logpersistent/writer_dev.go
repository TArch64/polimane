//go:build dev

package logpersistent

import (
	"io"
	"os"
)

type stdoutWriter struct {
	*bufferedWriter
	dest io.Writer
}

func newWriter(_ ProviderOptions) io.WriteCloser {
	writer := &stdoutWriter{
		dest: os.Stdout,
	}

	writer.bufferedWriter = newBufferedWriter(writer.putLogs)
	go writer.flushLoop()
	return writer
}

func (s *stdoutWriter) putLogs(rows [][]byte) {
	for _, row := range rows {
		_, _ = s.dest.Write(row)
	}
}
