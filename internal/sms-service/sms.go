package sms_service

import (
	"Diplom_Makarov/internal/utils"
	"sort"
	"strings"
)

// SMSData - Структура для хранения данных системы SMS
type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

// StartSmsService - Функция запускает сервис для получения данных о состоянии системы SMS из файла формата CSV.
// Данные считиваются и затем происходит их валидация и сортировка. Результат выполениния - [][]SMSData.
func StartSmsService() [][]SMSData {
	return SortedSmsData(ValidateSmsData(utils.ReadCsvFile(utils.ConfigData.SmsDataPath)))
}

// ValidateSmsData - Функция валидирует данные о состоянии системы SMS. На вход принимаем [][]string, результат
// выполнения - []SMSData
func ValidateSmsData(data [][]string) []SMSData {
	result := make([]SMSData, 0)

	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 4:
			continue
		case !utils.IsExist(utils.ConfigData.Alpha2Code, row[0]):
			continue
		case !utils.IsExist(utils.ConfigData.Providers, row[3]):
			continue
		default:
			var newSmsData SMSData
			newSmsData.Country = row[0]
			newSmsData.Bandwidth = row[1]
			newSmsData.ResponseTime = row[2]
			newSmsData.Provider = row[3]
			result = append(result, newSmsData)
		}
	}

	return result
}

// SortedSmsData - Функция сортирует данные о состоянии системы SMS. На вход принимаем []SMSData, результат
// выполнения -  срез [][]SMSData. Первый список отсортирован по названию провайдера от A до Z.
// Второй список отсортирован по названию страны от A до Z.
func SortedSmsData(sms []SMSData) [][]SMSData {
	result := make([][]SMSData, 0)
	smsDataSortedByCountryName := make([]SMSData, 0)
	smsDataSortedByProviderName := make([]SMSData, 0)

	for _, item := range sms {
		item.Country = utils.ConfigData.CountryAlpha2[item.Country]
		smsDataSortedByCountryName = append(smsDataSortedByCountryName, item)
		smsDataSortedByProviderName = append(smsDataSortedByProviderName, item)
	}

	sort.SliceStable(smsDataSortedByCountryName, func(i, j int) bool {
		return smsDataSortedByCountryName[i].Country < smsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(smsDataSortedByProviderName, func(i, j int) bool {
		return smsDataSortedByProviderName[i].Provider < smsDataSortedByProviderName[j].Provider
	})

	result = append(result, smsDataSortedByCountryName)
	result = append(result, smsDataSortedByProviderName)

	return result
}
