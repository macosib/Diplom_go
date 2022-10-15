package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

const ProvidersPath = "./internal/utils/allow_providers.csv"
const ProvidersCallPath = "./internal/utils/allow_providers_call.csv"
const AlphaCodesPath = "./internal/utils/countries_codes_and_coordinates.csv"

func ReadCsvFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetAlpha2Code(path string) []string {
	alphaCode := make([]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		alphaCode = append(alphaCode, line[1])
	}
	return alphaCode
}

func GetAllowProviders(path string) []string {
	allowProviders := make([]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		allowProviders = append(allowProviders, line[0])
	}
	return allowProviders
}

func IsExist(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func ToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	return number
}

func ToFloat32(str string) float32 {
	number, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Println(err)
	}
	return float32(number)
}
