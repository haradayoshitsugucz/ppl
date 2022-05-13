package dto

import (
	"net/http"
	"time"
)

type TimeContextKey struct{}

func GetTimeContextValue(r *http.Request) time.Time {
	return r.Context().Value(TimeContextKey{}).(time.Time)
}
