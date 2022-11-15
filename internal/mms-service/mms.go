package mms_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
)

// MMSData - Структура для хранения данных системы MMS
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// StartMmsService - Функция запускает сервис для получения данных о системе MMS. Результат выполения функциия -
// [][]MMSData, либо ошибка
func StartMmsService() ([][]MMSData, error) {
	data, err := getMmsData()

	if err != nil {
		var res [][]MMSData
		return res, err
	}

	return SortedMMSData(validateMmsData(data)), nil
}

// getMmsData - Функция отправляет запрос к API для получения данных о состоянии системы MMS.
// Результат выполения - []MMSData,nil, в случае если данные невозможно получить функция вернет - nil, error
func getMmsData() ([]MMSData, error) {
	var mmsData []MMSData

	response, err := http.Get(utils.ConfigData.MmsServicePath)
	if err != nil {
		return mmsData, errors.New("Не удалось отправить запрос к серверу о состоянии системы MMS")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return mmsData, errors.New("Ошибка получения данных с сервера о состоянии системы MMS")
	}

	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &mmsData); err != nil {
		return mmsData, errors.New("Ошибка при чтении данных с сервера о состоянии системы MMS")
	}

	return mmsData, nil
}

// validateMmsData - Функция валидирует данные о состоянии системы MMS. На вход принимаем []MMSData, результат
// выполнения - []MMSData.
func validateMmsData(data []MMSData) []MMSData {
	var result []MMSData

	for _, item := range data {
		if !utils.IsExist(utils.ConfigData.Alpha2Code, item.Country) || !utils.IsExist(utils.ConfigData.Providers, item.Provider) {
			continue
		}
		result = append(result, item)
	}

	return result
}

// SortedMMSData - Функция сортирует данные о состоянии системы MMS. На вход принимаем []MMSData, результат
// выполнения -  срез [][]MMSData. Первый список отсортирован по названию провайдера от A до Z.
// Второй список отсортирован по названию страны от A до Z.
func SortedMMSData(mms []MMSData) [][]MMSData {
	result := make([][]MMSData, 0)
	mmsDataSortedByCountryName := make([]MMSData, 0)
	mmsDataSortedByProviderName := make([]MMSData, 0)

	for _, item := range mms {
		item.Country = utils.ConfigData.CountryAlpha2[item.Country]
		mmsDataSortedByCountryName = append(mmsDataSortedByCountryName, item)
		mmsDataSortedByProviderName = append(mmsDataSortedByProviderName, item)
	}

	sort.SliceStable(mmsDataSortedByCountryName, func(i, j int) bool {
		return mmsDataSortedByCountryName[i].Country < mmsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(mmsDataSortedByProviderName, func(i, j int) bool {
		return mmsDataSortedByProviderName[i].Provider < mmsDataSortedByProviderName[j].Provider
	})

	result = append(result, mmsDataSortedByCountryName)
	result = append(result, mmsDataSortedByProviderName)

	return result
}
