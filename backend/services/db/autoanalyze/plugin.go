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
		logs:   newLogs(),
	}
}

func (p *Plugin) Name() string {
	return "auto_analyze"
}

func (p *Plugin) Initialize(db *gorm.DB) (err error) {
	err = db.
		Callback().
		Create().
		After("gorm:after_create").
		Register(p.Name()+":after_create", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	if err != nil {
		return err
	}

	err = db.
		Callback().
		Query().
		After("gorm:after_query").
		Register(p.Name()+":after_query", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	if err != nil {
		return err
	}

	err = db.
		Callback().
		Update().
		After("gorm:after_update").
		Register(p.Name()+":after_update", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	if err != nil {
		return err
	}

	err = db.
		Callback().
		Delete().
		After("gorm:after_delete").
		Register(p.Name()+":after_delete", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	if err != nil {
		return err
	}

	err = db.
		Callback().
		Row().
		After("gorm:after_row").
		Register(p.Name()+":after_row", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	err = db.
		Callback().
		Raw().
		After("gorm:after_raw").
		Register(p.Name()+":after_raw", func(query *gorm.DB) {
			p.onExecute(db, query)
		})

	return err
}

func (p *Plugin) onExecute(db, query *gorm.DB) {
	if p.rand.Float64() < 0.5 {
		return
	}

	queryStr := db.Explain(query.Statement.SQL.String(), query.Statement.Vars...)
	queryStr = strings.TrimSpace(queryStr)
	queryStr = strings.Trim(queryStr, "\n")

	var err error
	var explained string
	if explained, err = p.explainQuery(db, queryStr); err != nil {
		p.logErr("failed to explain query", err)
		return
	}

	if p.containsFullScan(explained) {
		p.logFullScan(queryStr, explained)
	}
}

func (p *Plugin) logErr(title string, err error) {
	p.stdout.Error(fmt.Sprintf("autoanalyze: %s", title),
		slog.String("err", err.Error()),
	)
}

func (p *Plugin) logFullScan(query, explained string) {
	p.logs.WriteLine("--- Full Scan Detected ---")
	p.logs.WriteLine(time.Now().Format(time.RFC3339))
	p.logs.WritePadding(1)
	p.logs.WriteLine(query)
	p.logs.WritePadding(1)
	p.logs.WriteLine(strings.ReplaceAll(explained, "FULL SCAN", fullScanMarker))
}
