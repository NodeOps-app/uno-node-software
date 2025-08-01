package time_utils

import (
	"math/big"
	"time"
)

// ConvertUnixToUTC converts a Unix timestamp in big.Int to a UTC time_utils with only year, month, and day.
func ConvertUnixToUTC(unixTime *big.Int) time.Time {
	// Convert big.Int to int64
	timestamp := unixTime.Int64()

	// Create time_utils.Time object from Unix timestamp
	t := time.Unix(timestamp, 0).UTC()

	// Return only the year, month, and day
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
