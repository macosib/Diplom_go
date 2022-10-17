package sms_service

import (
	"Diplom_Makarov/internal/utils"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func StartSmsService() [][]SMSData {
	path := "../simulator/skillbox-diploma/sms.data"
	return SortedSmsData(validateSmsData(utils.ReadCsvFile(path)))
}

func validateSmsData(data [][]string) []SMSData {
	codes := utils.GetAlpha2Code(utils.AlphaCodesPath)
	providers := utils.GetAllowProviders(utils.ProvidersPath)
	result := make([]SMSData, 0)
	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 4:
			continue
		case !utils.IsExist(codes, row[0]):
			continue
		case !utils.IsExist(providers, row[3]):
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

func SortedSmsData(sms []SMSData) [][]SMSData {
	countryArray := utils.GetCountryAlpha2Code(utils.AlphaCodesPath)
	result := make([][]SMSData, 0)
	smsDataSortedByCountryName := make([]SMSData, 0)
	smsDataSortedByProviderName := make([]SMSData, 0)
	for _, item := range sms {
		item.Country = countryArray[item.Country]
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
