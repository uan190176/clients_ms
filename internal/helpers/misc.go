package helpers

import (
	"strings"
)

// DeleteTabsAndNewLinesSymbols - delete tabs and new lines symbols
func DeleteTabsAndNewLinesSymbols(text string) string {
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	return text
}
