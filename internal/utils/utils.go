package utils

import (
	"encoding/csv"
	"log"
	"os"
)

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

func GetAlpha2Code() []string {
	alphaCode := make([]string, 0)
	data := ReadCsvFile("./internal/utils/countries_codes_and_coordinates.csv")
	for _, line := range data {
		alphaCode = append(alphaCode, line[1])
	}
	return alphaCode
}

func GetAllowProviders() []string {
	allowProviders := make([]string, 0)
	data := ReadCsvFile("./internal/utils/allow_providers.csv")
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
