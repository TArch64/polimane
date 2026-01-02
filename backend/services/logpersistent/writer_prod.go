//go:build !dev

package logpersistent

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"

	"polimane/backend/services/awscloudwatch"
	"polimane/backend/services/logstdout"
)

type logRow struct {
	Time      time.Time `json:"time"`
	Operation string    `json:"operation,omitempty"`
}

type cloudwatchWriter struct {
	*bufferedWriter
	ctx        context.Context
	cloudwatch *cloudwatchlogs.Client
	stdout     *logstdout.Logger
}

func newWriter(options ProviderOptions) io.WriteCloser {
	writer := &cloudwatchWriter{
		cloudwatch: options.Cloudwatch,
		ctx:        options.Ctx,
		stdout:     options.Stdout,
	}

	writer.bufferedWriter = newBufferedWriter(writer.putLogs)
	go writer.flushLoop()
	return writer
}

func (c *cloudwatchWriter) putLogs(rows [][]byte) {
	byStream := c.buildLogEvents(rows)

	for stream, events := range byStream {
		_, err := c.cloudwatch.PutLogEvents(c.ctx, &cloudwatchlogs.PutLogEventsInput{
			LogEvents:     events,
			LogGroupName:  &awscloudwatch.GroupPersistent,
			LogStreamName: &stream,
		})

		if err != nil {
			c.stdout.ErrorContext(c.ctx, "failed to put log events",
				slog.String("error", err.Error()),
				slog.String("stream", stream),
			)
			continue
		}
	}
}

func (c *cloudwatchWriter) buildLogEvents(rows [][]byte) map[string][]types.InputLogEvent {
	byStream := make(map[string][]types.InputLogEvent)

	for _, bytes := range rows {
		var row logRow
		if err := json.Unmarshal(bytes, &row); err != nil {
			c.stdout.ErrorContext(c.ctx, "failed to unmarshal log row",
				slog.String("error", err.Error()),
			)
			continue
		}

		if _, ok := byStream[row.Operation]; !ok {
			byStream[row.Operation] = []types.InputLogEvent{}
		}

		byStream[row.Operation] = append(byStream[row.Operation], types.InputLogEvent{
			Message:   aws.String(string(bytes)),
			Timestamp: aws.Int64(row.Time.UnixMilli()),
		})
	}

	return byStream
}
