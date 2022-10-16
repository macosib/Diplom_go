package support_service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func getSupportData() []SupportData {
	var supportData []SupportData
	response, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Printf("Ошибка получения данных")
		return supportData
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Printf("Ошибка получения данных")
		return supportData
	}
	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &supportData); err != nil {
		log.Printf("Ошибка при чтении данных")
		return supportData
	}
	return supportData
}

func StartSupportService() {
	data := getSupportData()
	fmt.Println(data)
}
