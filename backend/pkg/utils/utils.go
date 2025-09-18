package utils

import (
	"fmt"
	"time"
)

func TimestampToISO(timestamp int64) *string {
	if timestamp == 0 {
		return nil
	}
	t := time.Unix(timestamp, 0)
	iso := t.Format(time.RFC3339)
	return &iso
}

func ParseToRFC3339(dateStr string) *string {
	if dateStr == "" {
		return String(time.Now().UTC().Format(time.RFC3339))
	}

	// Possible layouts you may encounter
	layouts := []string{
		time.RFC3339,          // 2025-09-18T00:19:34Z or with offset
		"2006-01-02T15:04:05", // 2025-09-18T00:19:34 (no timezone)
		"2006-01-02 15:04:05", // 2025-09-18 00:19:34 (MySQL style)
		"2006-01-02",          // 2025-09-18 (date only)
	}

	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, dateStr); err == nil {
			return String(parsed.UTC().Format(time.RFC3339))
		}
	}

	// fallback if all parsing failed
	fmt.Printf("failed to parse date %q, falling back to current UTC time\n", dateStr)
	return String(time.Now().UTC().Format(time.RFC3339))
}
