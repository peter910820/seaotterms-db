package dto

type CommonResponse[T any] struct {
	StatusCode int    `json:"statusCode"` // http status code
	ErrMsg     string `json:"errMsg"`
	Data       T      `json:"data"`
}
