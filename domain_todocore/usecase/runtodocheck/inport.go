package runtodocheck

import (
	"github.com/KuraoHikari/gogen/domain_todocore/model/entity"
	"github.com/KuraoHikari/gogen/domain_todocore/model/vo"
	"github.com/KuraoHikari/gogen/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}
