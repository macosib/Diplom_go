package utils

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type PathConfig struct {
	Alpha2Code          []string
	Providers           []string
	ProvidersCall       []string
	ProvidersEmail      []string
	CountryAlpha2       map[string]string
	BillingDataPath     string
	EmailDataPath       string
	SmsDataPath         string
	VoiceDataPath       string
	SupportServicePath  string
	MmsServicePath      string
	IncidentServicePath string
}

func getBaseDir() string {
	baseDir, _ := os.Getwd()
	return strings.TrimRight(baseDir, "/tests")
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func newPathConfig() *PathConfig {
	return &PathConfig{
		GetAlpha2Code(path.Join(getBaseDir(), getEnv("ALPHA_CODES_PATH", ""))),
		GetAllowProviders(path.Join(getBaseDir(), getEnv("PROVIDERS_PATH", ""))),
		GetAllowProviders(path.Join(getBaseDir(), getEnv("PROVIDERS_CALL_PATH", ""))),
		GetAllowProviders(path.Join(getBaseDir(), getEnv("PROVIDERS_EMAIL_PATH", ""))),
		GetCountryAlpha2Code(path.Join(getBaseDir(), getEnv("ALPHA_CODES_PATH", ""))),
		path.Join(getBaseDir(), getEnv("BILLING_DATA_PATH", "")),
		path.Join(getBaseDir(), getEnv("EMAIL_DATA_PATH", "")),
		path.Join(getBaseDir(), getEnv("SMS_DATA_PATH", "")),
		path.Join(getBaseDir(), getEnv("VOICE_DATA_PATH", "")),
		getEnv("SUPPORT_SERVICE_PATH", ""),
		getEnv("MMS_SERVICE_PATH", ""),
		getEnv("INCIDENT_SERVICE_PATH", ""),
	}
}

var ConfigData = newPathConfig()

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

func GetCountryAlpha2Code(path string) map[string]string {
	result := make(map[string]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		result[line[1]] = line[0]
	}
	return result
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

func ConvertToBool(b byte) bool {
	if b == 48 {
		return false
	}
	return true
}

func ErrorToString(err ...error) string {
	var errorString string
	for _, item := range err {
		if item != nil {
			errorString += item.Error() + ", "
		}

	}
	return strings.TrimRight(errorString, ", ")
}
