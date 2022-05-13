// +build wireinject

package di

import (
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/controller"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/middleware"
	"github.com/haradayoshitsugucz/purple-server/api/service"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/infra/aws"
	"github.com/haradayoshitsugucz/purple-server/infra/dao"
)

// repository
var repositorySet = wire.NewSet(
	dao.NewDBCluster,
	dao.NewStatusRepository,
	dao.NewProductRepository,
	dao.NewBrandRepository,
	dao.NewTransactionRepository,
	aws.NewSecretsManager,
	aws.NewSecretsRepository,
)

// service
var serviceSet = wire.NewSet(
	service.NewStatusService,
	service.NewProductService,
)

// controller
var controllerSet = wire.NewSet(
	controller.NewStatusController,
	controller.NewProductController,
)

var middlewareSet = wire.NewSet(
	middleware.NewTimeMiddle,
)

func InitHandler(conf config.Config, timeNow func() time.Time, cli *http.Client) *ControllerAndMiddleware {
	wire.Build(
		repositorySet,
		serviceSet,
		controllerSet,
		middlewareSet,
		wire.Struct(new(controller.Controller), "*"),
		wire.Struct(new(middleware.Middleware), "*"),
		wire.Struct(new(ControllerAndMiddleware), "*"),
	)
	return &ControllerAndMiddleware{}
}
