package autoanalyze

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type pluginLogs struct {
	mx           *sync.Mutex
	fs           *os.Root
	logFile      *os.File
	logFileLimit int64
	logFileName  string
}

func newLogs() (logs *pluginLogs, err error) {
	logs = &pluginLogs{
		mx:           &sync.Mutex{},
		logFileName:  "auto_analyze",
		logFileLimit: 10 * 1024 * 1024, // 10 MB
	}

	logs.mx.Lock()
	defer logs.mx.Unlock()

	if err = os.MkdirAll("/tmp/app", 0755); err != nil {
		return nil, err
	}

	if logs.fs, err = os.OpenRoot("/tmp/app"); err != nil {
		return nil, err
	}

	if logs.logFile, err = logs.openFile(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (p *pluginLogs) Open(exec func(logs *pluginLogs) error) error {
	p.mx.Lock()
	defer p.mx.Unlock()

	info, err := p.logFile.Stat()
	if err != nil {
		return err
	}

	if info.Size() >= p.logFileLimit {
		err = p.forkFile()
	} else if info.Size() != 0 {
		err = p.WritePadding(2)
	}

	if err != nil {
		return err
	}

	return exec(p)
}

func (p *pluginLogs) WriteLine(text string) error {
	_, err := p.logFile.WriteString(text)
	return err
}

func (p *pluginLogs) WritePadding(n int) error {
	padding := strings.Repeat("\n", n*2)
	return p.WriteLine(padding)
}

func (p *pluginLogs) openFile() (*os.File, error) {
	return p.fs.OpenFile(p.logFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func (p *pluginLogs) forkFile() (err error) {
	if err = p.logFile.Close(); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	newName := fmt.Sprintf("%s_%s.log", p.logFileName, timestamp)
	if err = os.Rename(p.logFileName+".log", newName); err != nil {
		return err
	}

	p.logFile, err = p.openFile()
	return err
}
