package dto

type CommonRequest[T any] struct {
	Token string `json:"token"`
	Data  T      `json:"data"`
}
