package autoanalyze

import (
	"strings"
)

type pluginLogs struct{}

func newLogs() *pluginLogs {
	return &pluginLogs{}
}

func (p *pluginLogs) WriteLine(text string) {
	println(text)
}

func (p *pluginLogs) WritePadding(n int) {
	p.WriteLine(strings.Repeat("\n", n*2))
}
