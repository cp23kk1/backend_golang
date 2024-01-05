package common

type VocaVerseResponse struct {
	Status VocaVerseStatusResponse `json:"status"`
	Data   map[string]interface{}  `json:"data"`
}

type VocaVerseStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ConvertVocaVerseResponse(status VocaVerseStatusResponse, data map[string]interface{}) VocaVerseResponse {
	return VocaVerseResponse{
		Status: status,
		Data:   data,
	}
}
