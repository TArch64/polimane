package autoanalyze

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"

	"polimane/backend/base"
)

const (
	logFileLimit = 10 * 1024 * 1024 // 10 MB
	logFileName  = "auto_analyze"
)

var (
	fullScanMarker = base.Colored("FULL SCAN", base.AnsiRed)
)

type Plugin struct {
	mx      *sync.Mutex
	fs      *os.Root
	logFile *os.File
	rand    *rand.Rand
}

func New() gorm.Plugin {
	return &Plugin{
		mx:   &sync.Mutex{},
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *Plugin) Name() string {
	return "auto_analyze"
}

func (p *Plugin) Initialize(db *gorm.DB) (err error) {
	if err = os.MkdirAll("/tmp/app", 0755); err != nil {
		return err
	}

	if p.fs, err = os.OpenRoot("/tmp/app"); err != nil {
		return err
	}

	if p.logFile, err = p.openLogFile(logFileName); err != nil {
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

			var explained string
			if explained, err = p.explainQuery(db, query); err != nil {
				log.Println("autoanalyze: failed to analyze query:", err)
				return
			}

			if p.containsFullScan(explained) {
				if err = p.logFullScan(explained); err != nil {
					log.Println("autoanalyze: failed to log full scan:", err)
				}
			}
		})
}

func (p *Plugin) explainQuery(db, query *gorm.DB) (string, error) {
	explainSQL := "EXPLAIN ANALYZE " + db.Explain(query.Statement.SQL.String(), query.Statement.Vars...)

	var explained []string
	err := db.Raw(explainSQL).Scan(&explained).Error

	if err != nil {
		return "", err
	}

	return strings.Join(explained, "\n"), nil
}

func (p *Plugin) containsFullScan(explained string) bool {
	return strings.Contains(explained, "FULL SCAN")
}

func (p *Plugin) logFullScan(explained string) error {
	return p.writeLog(func(logs *os.File) (err error) {
		_, err = logs.WriteString("--- Full Scan Detected --- \n\n")
		if err != nil {
			return err
		}

		explained = strings.ReplaceAll(explained, "FULL SCAN", fullScanMarker)
		_, err = logs.WriteString(explained)
		return err
	})
}

func (p *Plugin) writeLog(exec func(logs *os.File) error) error {
	p.mx.Lock()
	defer p.mx.Unlock()

	info, err := p.logFile.Stat()
	if err != nil {
		return err
	}

	if info.Size() >= logFileLimit {
		if err = p.forkLogFile(); err != nil {
			return err
		}
	} else if info.Size() != 0 {
		if _, err = p.logFile.WriteString("\n\n\n\n"); err != nil {
			return err
		}
	}

	return exec(p.logFile)
}

func (p *Plugin) openLogFile(name string) (*os.File, error) {
	return p.fs.OpenFile(name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func (p *Plugin) forkLogFile() (err error) {
	if err = p.logFile.Close(); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	newName := fmt.Sprintf("%s_%s.log", logFileName, timestamp)
	if err = os.Rename(logFileName+".log", newName); err != nil {
		return err
	}

	p.logFile, err = p.openLogFile(logFileName)
	return err
}
