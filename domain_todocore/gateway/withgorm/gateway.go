package withgorm

import (
	"context"
	"fmt"

	"github.com/KuraoHikari/gogen/domain_todocore/model/entity"
	"github.com/KuraoHikari/gogen/domain_todocore/model/vo"
	"github.com/KuraoHikari/gogen/shared/gogen"
	"github.com/KuraoHikari/gogen/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen/shared/infrastructure/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	db		*gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "test")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entity.Todo{})

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	//   panic(err)
	// }
	// err = db.AutoMigrate(&entity.Todo{})
	// if err != nil {
	// 	panic(err)
	// }
	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db: 	db,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	var todoObjs []*entity.Todo
	var count int64

	err := r.db.Model(entity.Todo{}).Count(&count).Limit(size).Offset((page -1)*size).Find(&todoObjs).Error
	if err != nil {
		panic(err)
	}

	return todoObjs, count, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var todoObj entity.Todo

	err := r.db.First(&todoObj,"id = ?", todoID).Error
	if err != nil {
		panic(err)
	}
	return &todoObj, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")
	err := r.db.Save(obj).Error
	if err != nil {
		panic(err)
	}

	return nil
}
