package endpoints

import (
	"errors"
	"net/http"
	"reserva_vagas/internalerrors"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerErro(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		if err != nil {
			var statusCode int

			switch {
			case errors.Is(err, internalerrors.ErrInternal):
				statusCode = http.StatusInternalServerError // 500
			case errors.Is(err, internalerrors.ErrValidation):
				statusCode = http.StatusUnprocessableEntity // 422
			default:
				if status >= 400 && status < 600 {
					statusCode = status
				} else {
					statusCode = http.StatusBadRequest // 400
				}
			}

			render.Status(r, statusCode)
			render.JSON(w, r, map[string]string{"errors": err.Error()})
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
