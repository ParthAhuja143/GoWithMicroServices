package errors

import "net/http"

type HTTPError int

type Error struct {
	Message string `json:"message"`
	Description string `json:"description"`
	StatusCode int `json:"int"`
}

const (
	NoErr HTTPError = -1
	ErrProductNotFound HTTPError = 0	
	ErrInvalidURI      HTTPError = 1
	ErrUnmarshal       HTTPError = 2
	ErrInternalServer  HTTPError = 3
	ErrValidatingProduct HTTPError = 4
	ErrMarshal HTTPError = 5
)

func Err(e HTTPError) Error {
	response := Error{
		Message: e.Title(),
		Description: e.Description(),
		StatusCode: e.StatusCode(),
	}

	return response
}

func (e HTTPError) Title() string {
	switch e {
	case ErrProductNotFound:
		return "Product Not Found"
	case ErrInvalidURI:
		return "Invalid URI"
	case ErrUnmarshal:
		return "Error unmarshaling given json"
	case ErrInternalServer:
		return "Internal Server Error"
	case ErrValidatingProduct:
		return "Error Validing Product"
	case ErrMarshal:
		return "Error marshaling given json"
	default:
		return ""
	}
}

func (e HTTPError) Description() string {
	switch e {
	case ErrProductNotFound:
		return "Product Not Found"
	case ErrInvalidURI:
		return "Invalid URI"
	case ErrUnmarshal:
		return "Error unmarshaling given json"
	case ErrInternalServer:
		return "Internal Server Error"
	case ErrMarshal:
		return "Error marshaling given json"
	case ErrValidatingProduct:
		return "Error Validating Product"
	default:
		return ""
	}
}

func (e HTTPError) StatusCode() int {
	switch e {
	case ErrInvalidURI,
		ErrUnmarshal,
		ErrMarshal,
		ErrValidatingProduct:
		return http.StatusBadRequest
	case ErrProductNotFound:
		return http.StatusNotFound
	case ErrInternalServer:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
