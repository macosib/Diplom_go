package mms_service

import (
	"Diplom_Makarov/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
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

func StartMmsService() ([][]MMSData, error) {
	data, err := getMmsData()
	if err != nil {
		var res [][]MMSData
		return res, err
	}
	return SortedMMSData(validateMmsData(data)), nil
}

func SortedMMSData(mms []MMSData) [][]MMSData {
	countryArray := utils.GetCountryAlpha2Code(utils.AlphaCodesPath)
	result := make([][]MMSData, 0)
	mmsDataSortedByCountryName := make([]MMSData, 0)
	mmsDataSortedByProviderName := make([]MMSData, 0)
	for _, item := range mms {
		item.Country = countryArray[item.Country]
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
