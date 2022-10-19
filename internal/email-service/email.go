package email_service

import (
	"Diplom_Makarov/internal/utils"
	"sort"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func StartEmailService() map[string][][]EmailData {
	path := "../simulator/skillbox-diploma/email.data"
	data := validateEmailData(utils.ReadCsvFile(path))
	return SortedEmailData(data)
}

func validateEmailData(data [][]string) []EmailData {
	result := make([]EmailData, 0)
	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 3:
			continue
		case !utils.IsExist(utils.ConfigData.Alpha2Code, row[0]):
			continue
		case !utils.IsExist(utils.ConfigData.ProvidersEmail, row[1]):
			continue
		default:
			var newEmailData EmailData
			newEmailData.Country = row[0]
			newEmailData.Provider = row[1]
			newEmailData.DeliveryTime = utils.ToInt(row[2])
			result = append(result, newEmailData)
		}
	}
	return result
}

func SortedEmailData(emailData []EmailData) map[string][][]EmailData {
	result := make(map[string][][]EmailData)
	sort.SliceStable(emailData, func(i, j int) bool {
		return emailData[i].Country > emailData[j].Country
	})
	uniqueCountryList := getCountry(emailData)
	for _, country := range uniqueCountryList {
		res := make([]EmailData, 0)
		for _, data := range emailData {
			if country == data.Country {
				res = append(res, data)
			}
		}
		sort.SliceStable(res, func(i, j int) bool {
			return res[i].DeliveryTime > res[j].DeliveryTime
		})
		result[country] = [][]EmailData{res[:3], res[len(res)-3:]}
	}
	return result
}

func getCountry(data []EmailData) []string {
	result := make([]string, 0)

	for _, item := range data {
		result = append(result, item.Country)
	}

	return uniqueCountry(result)
}

func uniqueCountry(array []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0)

	for _, item := range array {
		if _, value := keys[item]; !value {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
