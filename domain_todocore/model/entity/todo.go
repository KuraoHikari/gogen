package entity

import (
	"strings"
	"time"

	"github.com/KuraoHikari/gogen/domain_todocore/model/errorenum"
	"github.com/KuraoHikari/gogen/domain_todocore/model/vo"
)

type Todo struct {
	ID      	vo.TodoID 	`bson:"_id" json:"id"`
	Created 	time.Time 	`bson:"created" json:"created"`
	Updated 	time.Time 	`bson:"updated" json:"updated"`
	Message 	string 		`json:"message"`
	Checked		bool		`json:"checked"`

}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Message 	string 		`json:"message"`
}

func (r TodoCreateRequest) Validate() error {

	if strings.TrimSpace(r.Message) == ""{
		return errorenum.MessageMustNotEmpty
	}

	return nil
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewTodoID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj Todo
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now
	obj.Message = req.Message
	obj.Checked = false

	return &obj, nil
}

type TodoUpdateRequest struct {
	Now          time.Time `json:"-"`
	// add field to update here ...
}

func (r TodoUpdateRequest) Validate() error {
	return nil
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	// add validation and assignment value here ...
	err := req.Validate()
	if err != nil {
		return  err
	}
	r.Updated = req.Now

	return nil
}

func (r *Todo) Check() error{

	if r.Checked {
		return errorenum.TodoHasBeenChecked
	}
	r.Checked = true

	return nil
}
