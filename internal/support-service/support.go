package support_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// SupportData - Структура для хранения данных системы Support
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// StartSupportService - Функция запускает сервис для получения данных о системе Support. Результат выполения функциия -
// []SupportData, либо ошибка.
func StartSupportService() ([]int, error) {
	data, err := getSupportData()
	if err != nil {
		return []int{0, 0}, err
	}
	return validSupportData(data), nil
}

// getSupportData - Функция отправляет запрос к API для получения данных о текущей загрузке команды  службы поддержки.
// Результат выполения - []SupportData,nil, в случае если данные невозможно получить функция вернет - nil, error
func getSupportData() ([]SupportData, error) {
	var supportData []SupportData

	response, err := http.Get(utils.ConfigData.SupportServicePath)

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

// validSupportData - Функция валидирует данные о состоянии системы Support. На вход принимаем []SupportData, результат
// выполнения - []int. Срез из двух int, первый из которых показывает загруженность службы поддержки (1–3),
// а второй — среднее время ожидания ответа.
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
