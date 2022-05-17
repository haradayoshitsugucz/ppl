package util

import (
	"os"
	"time"
)

func StringToPointer(s string) *string {
	return &s
}

func Int64ToPointer(i int64) *int64 {
	return &i
}

func TimeToPointer(t time.Time) *time.Time {
	return &t
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
