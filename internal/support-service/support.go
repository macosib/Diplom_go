package support_service

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func getSupportData() ([]SupportData, error) {
	var supportData []SupportData
	response, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return supportData, errors.New("Ошибка при запросе данных с сервера")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return supportData, errors.New("Ошибка при получении данных с сервера")
	}
	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &supportData); err != nil {
		return supportData, errors.New("Ошибка при чтении данных с сервера")
	}
	return supportData, nil
}

func StartSupportService() ([]SupportData, error) {
	data, err := getSupportData()
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}
	return data, nil
}
