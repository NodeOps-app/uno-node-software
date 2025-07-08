package common

import (
	"time"
)

type BackoffHandler struct {
	initial time.Duration
	max     time.Duration
	current time.Duration
}

func NewBackoffHandler(initial, max time.Duration) *BackoffHandler {
	return &BackoffHandler{
		initial: initial,
		max:     max,
		current: initial,
	}
}

func (b *BackoffHandler) Next() time.Duration {
	b.current *= 2
	if b.current > b.max {
		b.current = b.max
	}
	return b.current
}

func (b *BackoffHandler) Reset() {
	b.current = b.initial
}

func (b *BackoffHandler) LimitReached() bool {
	return b.current > b.max
}
