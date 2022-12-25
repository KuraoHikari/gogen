package runtodocreate

import (
	"github.com/KuraoHikari/gogen/domain_todocore/model/entity"
	"github.com/KuraoHikari/gogen/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
