//go:build dev

package awsconfig

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/smithy-go/logging"

	"polimane/backend/services/logstdout"
)

func configure(options *Options, loadOptions *config.LoadOptions) {
	loadOptions.DefaultRegion = options.Env.AWS.Region
	loadOptions.Credentials = credentials.NewStaticCredentialsProvider(
		options.Env.AWS.AccessKeyID,
		options.Env.AWS.SecretAccessKey,
		"",
	)

	clientLogMode := aws.LogRetries | aws.LogRequest
	loadOptions.ClientLogMode = &clientLogMode

	loadOptions.Logger = &stdoutAdapter{
		stdout: options.Stdout,
		ctx:    context.Background(),
	}
}

type stdoutAdapter struct {
	stdout *logstdout.Logger
	ctx    context.Context
}

func (s *stdoutAdapter) WithContext(ctx context.Context) logging.Logger {
	return &stdoutAdapter{
		stdout: s.stdout,
		ctx:    ctx,
	}
}

func (s *stdoutAdapter) Logf(classification logging.Classification, format string, v ...interface{}) {
	rawReq := fmt.Sprintf(format, v...)
	rawReq = strings.TrimPrefix(rawReq, "Request\n")

	req, err := s.parseRequest(rawReq)
	if err != nil {
		s.stdout.ErrorContext(s.ctx, "AWS LOGGER: failed to parse HTTP request",
			slog.String("error", err.Error()),
			slog.String("raw_request", rawReq),
		)
		return
	}

	headersAttrs := make([]any, 0, len(req.Header))
	for key, values := range req.Header {
		headersAttrs = append(headersAttrs, slog.String(key, strings.Join(values, ", ")))
	}

	level := s.getLogLevel(classification)

	s.stdout.Log(s.ctx, level, "AWS Request",
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Group("headers", headersAttrs...),
	)
}

func (s *stdoutAdapter) getLogLevel(classification logging.Classification) slog.Level {
	if classification == logging.Warn {
		return slog.LevelWarn
	}
	return slog.LevelInfo
}

func (s *stdoutAdapter) parseRequest(rawReq string) (*http.Request, error) {
	reader := bufio.NewReader(strings.NewReader(rawReq))
	return http.ReadRequest(reader)
}
