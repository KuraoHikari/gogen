package resapi

import (
	"github.com/KuraoHikari/gogen/shared/gogen"
	"github.com/KuraoHikari/gogen/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen/shared/infrastructure/logger"
	"github.com/KuraoHikari/gogen/shared/infrastructure/token"
	"github.com/gin-gonic/gin"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	resource := router.Group("/api/v1", r.authentication())
	resource.GET("/todo", r.authorization(), r.getAllTodoHandler())
	resource.PUT("/todo/:todo_id", r.authorization(), r.runTodoCheckHandler())
	resource.POST("/POST", r.authorization(), r.runTodoCreateHandler())

}
