package util

import "time"

func StringToPointer(s string) *string {
	return &s
}

func Int64ToPointer(i int64) *int64 {
	return &i
}

func TimeToPointer(t time.Time) *time.Time {
	return &t
}
