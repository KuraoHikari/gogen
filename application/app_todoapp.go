package application

import (
	"github.com/KuraoHikari/gogen/domain_todocore/controller/resapi"
	"github.com/KuraoHikari/gogen/domain_todocore/gateway/withgorm"
	"github.com/KuraoHikari/gogen/domain_todocore/usecase/getalltodo"
	"github.com/KuraoHikari/gogen/domain_todocore/usecase/runtodocheck"
	"github.com/KuraoHikari/gogen/domain_todocore/usecase/runtodocreate"
	"github.com/KuraoHikari/gogen/shared/gogen"
	"github.com/KuraoHikari/gogen/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen/shared/infrastructure/logger"
	"github.com/KuraoHikari/gogen/shared/infrastructure/server"
	"github.com/KuraoHikari/gogen/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := resapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
