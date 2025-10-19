package base

import "fmt"

const (
	AnsiRed     = "\033[31m"
	AnsiGreen   = "\033[32m"
	AnsiYellow  = "\033[33m"
	AnsiBlue    = "\033[34m"
	AnsiMagenta = "\033[35m"
	AnsiCyan    = "\033[36m"
	AnsiReset   = "\033[0m"
)

func Colored(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, AnsiReset)
}
