package runtodocreate

import "github.com/KuraoHikari/gogen/domain_todocore/model/repository"

type Outport interface {
	repository.SaveTodoRepo
}
