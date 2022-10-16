package incident_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func getIncidentData() ([]IncidentData, error) {
	var incidentData []IncidentData
	response, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		return incidentData, errors.New("Ошибка при запросе к серверу")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return incidentData, errors.New("Ошибка получения данных с сервера")
	}
	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &incidentData); err != nil {
		return incidentData, errors.New("Ошибка при чтении данных с сервера")
	}
	return incidentData, nil
}

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

func StartIncidentService() ([]IncidentData, error) {
	data, err := getIncidentData()
	if err != nil {
		return data, err
	}
	return validateIncidentData(data), nil
}
