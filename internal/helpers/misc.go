package helpers

import (
	"fmt"
	"strings"
	"time"
)

// GetDeltaDatesInHours - returns delta in hours between 2 dates
func GetDeltaDatesInHours(d1 string, d2 time.Time) int {
	dd1, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v", d1))
	if err != nil {
		return -1
	}

	d := d2.Sub(dd1)
	fmt.Println(d)
	return int(d2.Sub(dd1) / time.Hour)
}

// DeleteTabsAndNewLinesSymbols - delete tabs and new lines symbols
func DeleteTabsAndNewLinesSymbols(text string) string {
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	return text
}
