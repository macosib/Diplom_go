package support_service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func StartSupportService() ([]int, error) {
	data, err := getSupportData()
	if err != nil {
		return []int{0, 0}, err
	}
	return validSupportData(data), nil
}

func getSupportData() ([]SupportData, error) {
	var supportData []SupportData

	response, err := http.Get("http://127.0.0.1:8383/support")

	if err != nil {
		return supportData, errors.New("Не удалось отправить запрос к серверу о системе Support")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return supportData, errors.New("Ошибка при получении данных с сервера о системе Support")
	}

	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &supportData); err != nil {
		return supportData, errors.New("Ошибка при чтении данных с сервера о системе Support")
	}

	return supportData, nil
}

func validSupportData(data []SupportData) []int {
	result := make([]int, 0)
	var totalTopic, load, averageTime int

	for _, item := range data {
		totalTopic += item.ActiveTickets
	}

	switch {
	case totalTopic < 9:
		load = 1
	case totalTopic <= 16:
		load = 2
	default:
		load = 3
	}

	averageTime = int((float64(60) / float64(18)) * float64(totalTopic))
	result = append(result, load)
	result = append(result, averageTime)

	return result

}
