package entity

type TaskDTOInput struct {
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}

type TaskDTOOutput struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}
