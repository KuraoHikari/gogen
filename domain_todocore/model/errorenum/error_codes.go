package errorenum

import "github.com/KuraoHikari/gogen/shared/model/apperror"

const (
	SomethingError      apperror.ErrorType = "ER0000 something error"
	MessageMustNotEmpty apperror.ErrorType = "ER0001 message must not empty"
	TodoHasBeenChecked  apperror.ErrorType = "ER0002 todo has been checked"
)
