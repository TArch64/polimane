package autoanalyze

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"

	"polimane/backend/base"
	"polimane/backend/services/logstdout"
)

var (
	fullScanMarker = base.Colored("FULL SCAN", base.AnsiRed)
)

type Plugin struct {
	logs   *pluginLogs
	rand   *rand.Rand
	stdout *logstdout.Logger
}

type PluginOptions struct {
	Stdout *logstdout.Logger
}

func New(options *PluginOptions) gorm.Plugin {
	return &Plugin{
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		stdout: options.Stdout,
	}
}

func (p *Plugin) Name() string {
	return "auto_analyze"
}

func (p *Plugin) Initialize(db *gorm.DB) (err error) {
	if p.logs, err = newLogs(); err != nil {
		return err
	}

	return db.
		Callback().
		Query().
		After("gorm:query").
		Register(p.Name()+":after_query", func(query *gorm.DB) {
			if p.rand.Float64() < 0.5 {
				return
			}

			queryStr := db.Explain(query.Statement.SQL.String(), query.Statement.Vars...)

			var explained string
			if explained, err = p.explainQuery(db, queryStr); err != nil {
				p.logErr("failed to explain query", err)
				return
			}

			if p.containsFullScan(explained) {
				if err = p.logFullScan(queryStr, explained); err != nil {
					p.logErr("failed to log full scan", err)
				}
			}
		})
}

func (p *Plugin) logErr(title string, err error) {
	p.stdout.Error(fmt.Sprintf("autoanalyze: %s", title),
		slog.String("err", err.Error()),
	)
}

func (p *Plugin) logFullScan(query, explained string) error {
	return p.logs.Open(func(logs *pluginLogs) (err error) {
		if err = logs.WriteLine("--- Full Scan Detected ---"); err != nil {
			return err
		}
		if err = logs.WriteLine(time.Now().Format(time.RFC3339)); err != nil {
			return err
		}
		if err = logs.WritePadding(1); err != nil {
			return err
		}
		if err = logs.WriteLine(query); err != nil {
			return err
		}
		if err = logs.WritePadding(1); err != nil {
			return err
		}
		explained = strings.ReplaceAll(explained, "FULL SCAN", fullScanMarker)
		return logs.WriteLine(explained)
	})
}
