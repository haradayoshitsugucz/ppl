package controller

import (
	"fmt"
	"net/http"

	"github.com/haradayoshitsugucz/purple-server/api/domain/dto"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/response"
	"github.com/haradayoshitsugucz/purple-server/api/service"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/constant"
	"github.com/haradayoshitsugucz/purple-server/logger"
)

type StatusController interface {
	Ping() http.HandlerFunc
}

type StatusControllerImpl struct {
	conf          config.Config
	statusService service.StatusService
}

func NewStatusController(
	conf config.Config,
	statusService service.StatusService) StatusController {
	return &StatusControllerImpl{
		conf:          conf,
		statusService: statusService,
	}
}

func (c *StatusControllerImpl) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		now := dto.GetTimeContextValue(r)
		logger.GetLogger().Debug(fmt.Sprintf("%s", now.Format(constant.DatetimeFormat)))

		err := c.statusService.Ping()
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}
		response.SuccessWith(w, map[string]int{"status": 200})
		return
	}
}
