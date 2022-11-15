package email_service

import (
	"Diplom_Makarov/internal/utils"
	"sort"
	"strings"
)

// EmailData - Структура для хранения данных системы Email
type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

// StartEmailService - Функция запускает сервис для получения данных о состоянии системы Email из файла формата CSV.
// Данные считиваются и затем происходит их валидация и сортировка. Результат выполениния - map[string][][]EmailData.
func StartEmailService() map[string][][]EmailData {
	return SortedEmailData(validateEmailData(utils.ReadCsvFile(utils.ConfigData.EmailDataPath)))
}

// validateEmailData - Функция валидирует данные о состоянии системы Email. На вход принимаем [][]string, результат
// выполнения - []EmailData
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

// SortedEmailData - Функция сортирует данные о состоянии системы Email. На вход принимаем []EmailData, результат
// выполнения -  map[string][][]EmailData.
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

// getCountry - Функция получает список стран из среза []EmailData. На вход принимаем []EmailData, результат
// выполнения -  []string.
func getCountry(data []EmailData) []string {
	result := make([]string, 0)

	for _, item := range data {
		result = append(result, item.Country)
	}

	return uniqueCountry(result)
}

// uniqueCountry - Функция удаляет все повторяющиеся элементы. а вход принимаем []string, результат
// выполнения -  []string.
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
