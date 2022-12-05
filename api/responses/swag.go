package responses

// StatusOk 200
type StatusOk struct {
	Status string `json:"status" example:"ok"`
}

// ErrorUnauthorized 401
type ErrorUnauthorized struct {
	Error string `json:"errors" example:"Unauthorized"`
}

// ErrorNotFound 404
type ErrorNotFound struct {
	Error string `json:"errors" example:"NotFound"`
}

// ErrorInternalServerError 500
type ErrorInternalServerError struct {
	Error string `json:"errors" example:"InternalServerError"`
}
