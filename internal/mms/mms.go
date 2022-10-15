package mms

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func getMmsData() ([]MMSData, error) {
	response, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil, errors.New("Ошибка получения данных")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var mmsData []MMSData
	if err := json.Unmarshal(body, &mmsData); err != nil {
		return nil, errors.New("Ошибка при чтении данных")
	}
	return mmsData, nil
}

func validateMmsData(data []MMSData) []MMSData {
	code := utils.GetAlpha2Code()
	providers := utils.GetAllowProviders()
	var result []MMSData
	for _, item := range data {
		if !utils.IsExist(code, item.Country) || !utils.IsExist(providers, item.Provider) {
			continue
		}
		result = append(result, item)
	}
	return result
}

func StartMmsService() {
	data, err := getMmsData()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println(validateMmsData(data))
}
