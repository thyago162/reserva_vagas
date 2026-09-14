package internalerrors

import "errors"

var ErrInternal error = errors.New("Internal server error")
var ErrRegistroExiste error = errors.New("Registro já existe")
