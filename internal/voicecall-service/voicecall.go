package voicecall_service

import (
	"Diplom_Makarov/internal/utils"
	"log"
	"strings"
)

type VoiceData struct {
	Country             string
	CurrentLoad         int
	AverageResponseTime int
	Provider            string
	ConnectionStability float32
	UnknownParameter    int
	TTFB                int
	MedianCallDuration  int
}

func StartVoiceService() {
	res := validateVoiceData(utils.ReadCsvFile("../simulator/skillbox-diploma/voice.data"))
	log.Println(res)

}

func validateVoiceData(data [][]string) []VoiceData {
	codes := utils.GetAlpha2Code(utils.AlphaCodesPath)
	providers := utils.GetAllowProviders(utils.ProvidersCallPath)
	result := make([]VoiceData, 0)
	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 8:
			continue
		case !utils.IsExist(codes, row[0]):
			continue
		case !utils.IsExist(providers, row[3]):
			continue
		default:
			var newVoiceData VoiceData
			newVoiceData.Country = row[0]
			newVoiceData.CurrentLoad = utils.ToInt(row[1])
			newVoiceData.AverageResponseTime = utils.ToInt(row[2])
			newVoiceData.Provider = row[3]
			newVoiceData.ConnectionStability = utils.ToFloat32(row[4])
			newVoiceData.UnknownParameter = utils.ToInt(row[5])
			newVoiceData.TTFB = utils.ToInt(row[6])
			newVoiceData.MedianCallDuration = utils.ToInt(row[7])
			result = append(result, newVoiceData)
		}
	}
	return result
}
