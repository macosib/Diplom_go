package sms_service

import (
	"Diplom_Makarov/internal/utils"
	"log"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func StartSmsService() {
	res := validateSmsData(utils.ReadCsvFile("../simulator/skillbox-diploma/sms.data"))
	log.Println(res)

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
			newSmsData.Provider = row[2]
			result = append(result, newSmsData)
		}
	}
	return result
}
