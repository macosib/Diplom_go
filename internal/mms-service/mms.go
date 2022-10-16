package mms_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
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
	var mmsData []MMSData
	response, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return mmsData, errors.New("Ошибка при запросе к серверу")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return mmsData, errors.New("Ошибка получения данных с сервера")
	}
	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &mmsData); err != nil {
		return mmsData, errors.New("Ошибка при чтении данных с сервера")
	}
	return mmsData, nil
}

func validateMmsData(data []MMSData) []MMSData {
	codes := utils.GetAlpha2Code(utils.AlphaCodesPath)
	providers := utils.GetAllowProviders(utils.ProvidersPath)
	var result []MMSData
	for _, item := range data {
		if !utils.IsExist(codes, item.Country) || !utils.IsExist(providers, item.Provider) {
			continue
		}
		result = append(result, item)
	}
	return result
}

func StartMmsService() ([]MMSData, error) {
	data, err := getMmsData()
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}
	return validateMmsData(data), nil
}
