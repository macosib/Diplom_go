package incident_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func getIncidentData() []IncidentData {
	var incidentData []IncidentData
	response, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Printf("Ошибка получения данных")
		return incidentData
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Printf("Ошибка получения данных")
		return incidentData
	}
	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &incidentData); err != nil {
		log.Printf("Ошибка при чтении данных")
		return incidentData
	}
	return incidentData
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

func StartIncidentService() {
	data := getIncidentData()
	fmt.Println(validateIncidentData(data))
}
