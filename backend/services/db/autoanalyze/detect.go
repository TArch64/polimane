package autoanalyze

import (
	"strings"

	"gorm.io/gorm"
)

func (p *Plugin) explainQuery(db *gorm.DB, query string) (string, error) {
	var explained []string

	err := db.
		Raw("EXPLAIN " + query).
		Scan(&explained).
		Error

	if err != nil {
		return "", err
	}

	return strings.Join(explained, "\n"), nil
}

func (p *Plugin) containsFullScan(explained string) bool {
	return strings.Contains(explained, "FULL SCAN")
}
