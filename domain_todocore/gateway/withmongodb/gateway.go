package withmongodb

import (
	"context"

	"github.com/KuraoHikari/gogen/domain_todocore/model/entity"
	"github.com/KuraoHikari/gogen/domain_todocore/model/vo"
	"github.com/KuraoHikari/gogen/shared/gogen"
	"github.com/KuraoHikari/gogen/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen/shared/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	client  *mongo.Client
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {
	uri := "mongodb+srv://KuraoHikari:GangSetan12345@cluster0.qoyvi.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}


	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		client: client,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")
	coll := r.client.Database("test").Collection("todo")
	filter := bson.D{{}}
	skip := int64(size * (page-1))
	limit := int64(size)
	countOpts :=options.CountOptions{
		Limit: &limit,
		Skip: &skip,
	}
	count, err :=coll.CountDocuments(context.TODO(),filter, &countOpts)
	if err != nil {
		return nil, 0,err
	}
	findOpts :=options.FindOptions{
		Limit: &limit,
		Skip: &skip,
	}
	cursor, err := coll.Find(context.TODO(), filter, &findOpts)
	if err != nil {
		return nil, 0,err
	}
	res := make([]*entity.Todo,0)
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil, 0,err
	}
	return res, count, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	coll := r.client.Database("test").Collection("todo")
	filter := bson.D{{}}
	var result entity.Todo
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				return nil, nil
			}
			return nil, err
		}

	return &result, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	coll := r.client.Database("test").Collection("todo")

	result, err := coll.InsertOne(context.TODO(), obj)
	if err != nil {
		panic(err)
	}
	r.log.Info(ctx, "inserted with id %v", result.InsertedID)
	return nil
}
