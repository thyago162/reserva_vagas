package internalerrors

import "errors"

var ErrInternal error = errors.New("Internal server error")
var ErrRegistroExiste error = errors.New("Registro já existe")
var ErrVagasIndisponiveis error = errors.New("Não há vagas disponíveis")
var ErrVagasReservadasIndisponiveis = errors.New("Não existem vagas reservadas para decrementar")
var ErrTransacaoInvalida = errors.New("Transação inválida ou tipo incorreto")
