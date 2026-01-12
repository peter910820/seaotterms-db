package blog

type GetSystemTodo struct {
	ID         *int    `json:"id"`
	SystemName *string `json:"systemName"`
	Status     *int    `json:"status"`
}
