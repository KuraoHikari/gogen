package withgorm

import (
	"context"

	"github.com/KuraoHikari/gogen/domain_todocore/model/entity"
	"github.com/KuraoHikari/gogen/domain_todocore/model/vo"
	"github.com/KuraoHikari/gogen/shared/gogen"
	"github.com/KuraoHikari/gogen/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	return nil
}
