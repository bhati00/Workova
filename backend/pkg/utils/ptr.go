package utils

func String(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}
