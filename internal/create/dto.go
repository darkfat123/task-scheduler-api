package create

type CreateRequest struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	FrequencyDate string `json:"frequency_date"`
	FrequencyTime string `json:"frequency_time"`
	MaxRetries    int32  `json:"max_retries"`
}

type CreateResponse struct {
	Message string `json:"message"`
}
