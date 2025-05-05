package getall

type GetAllResponse struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	FrequencyDate string `json:"frequency_date"`
	FrequencyTime string `json:"frequency_time"`
	Status        string `json:"status"`
	IsEnabled     bool   `json:"is_enabled"`
}
