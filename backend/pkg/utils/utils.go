package utils

import "time"

func TimestampToISO(timestamp int64) *string {
	if timestamp == 0 {
		return nil
	}
	t := time.Unix(timestamp, 0)
	iso := t.Format(time.RFC3339)
	return &iso
}
