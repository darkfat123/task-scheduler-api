package get

import "time"

type GetByCodeResponse struct {
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	FrequencyDate string    `json:"frequency_date"`
	FrequencyTime string    `json:"frequency_time"`
	NextRunAt     time.Time `json:"next_run_at"`
	LastRunAt     time.Time `json:"last_run_at"`
	MaxRetries    int32     `json:"max_retries"`
	Status        string    `json:"status"`
	IsEnabled     bool      `json:"is_enabled"`
}

type GetAllResponse struct {
	Message string `json:"message"`
}
