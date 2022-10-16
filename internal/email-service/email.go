package email_service

import (
	"Diplom_Makarov/internal/utils"
	"log"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func StartEmailService() {
	res := validateEmailData(utils.ReadCsvFile("../simulator/skillbox-diploma/email.data"))
	log.Println(res)

}

func validateEmailData(data [][]string) []EmailData {
	codes := utils.GetAlpha2Code(utils.AlphaCodesPath)
	providers := utils.GetAllowProviders(utils.ProvidersEmailPath)
	result := make([]EmailData, 0)
	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 3:
			continue
		case !utils.IsExist(codes, row[0]):
			continue
		case !utils.IsExist(providers, row[1]):
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
