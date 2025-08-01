package constants

import "time"

const (
	DefaultPacemakerInterval = 30 * time.Second
	DefaultPollingInterval   = 1 * time.Minute
	DefaultBackoffInitial    = 1 * time.Second
	DefaultBackoffMax        = 15 * time.Second
	DefaultConnectionTimeout = 5 * time.Second
)
