package jobs

import (
	letslog "github.com/letsgo-framework/letsgo-mux/log"
)

// Greet logs Hello Jobs
func Greet() {
	letslog.Debug("Hello Jobs")
}
