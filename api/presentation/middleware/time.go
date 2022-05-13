package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/haradayoshitsugucz/purple-server/api/domain/dto"
	"github.com/haradayoshitsugucz/purple-server/logger"
)

type TimeMiddle interface {
	Now() func(next http.Handler) http.Handler
}

type TimeMiddleImpl struct {
	timeNow func() time.Time
}

func NewTimeMiddle(timeNow func() time.Time) TimeMiddle {
	return &TimeMiddleImpl{
		timeNow: timeNow,
	}
}

func (m *TimeMiddleImpl) Now() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			logger.GetLogger().Debug(fmt.Sprintf("[TimeMiddle] get time (system_date=%s)", m.timeNow().Format(time.RFC3339)))

			// 現在時刻をcontextに保存
			context := context.WithValue(r.Context(), dto.TimeContextKey{}, m.timeNow())

			next.ServeHTTP(w, r.WithContext(context))
		})
	}
}
