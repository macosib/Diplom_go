package incident_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
)

// IncidentData - Структура для хранения данных о системе истории инцидентов
type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// StartIncidentService - Функция запускает сервис для получения данных о системе истории инцидентов. Результат выполения функциия -
// []IncidentData, либо ошибка
func StartIncidentService() ([]IncidentData, error) {
	data, err := getIncidentData()
	if err != nil {
		return data, err
	}
	return validateIncidentData(data), nil
}

// getIncidentData - Функция отправляет запрос к API для получения данных о системе истории инцидентов.
// Результат выполения - []IncidentData,nil, в случае если данные невозможно получить функция вернет - nil, error
func getIncidentData() ([]IncidentData, error) {
	var incidentData []IncidentData

	response, err := http.Get(utils.ConfigData.IncidentServicePath)
	if err != nil {
		return incidentData, errors.New("Не удалось отправить запрос к серверу о системе истории инцидентов")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return incidentData, errors.New("Ошибка получения данных с сервера о системе истории инцидентов")
	}

	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &incidentData); err != nil {
		return incidentData, errors.New("Ошибка при чтении данных с сервера о системе истории инцидентов")
	}

	sort.SliceStable(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})

	return incidentData, nil
}

// validateIncidentData - Функция валидирует данные о системе истории инцидентов. На вход принимаем []IncidentData, результат
// выполнения - []IncidentData.
func validateIncidentData(data []IncidentData) []IncidentData {
	statusAllow := [2]string{"active", "closed"}
	var result []IncidentData
	for _, item := range data {
		if !utils.IsExist(statusAllow[:], item.Status) {
			continue
		}
		result = append(result, item)
	}
	return result
}
