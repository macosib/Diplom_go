package utils

import (
	"encoding/csv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var ConfigData = newPathConfig()

// PathConfig - структура для хранения параметров конфигурации путей.
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

// getBaseDir - Функция для получения корневого пути проекта. Возвращает string
func getBaseDir() string {
	baseDir, _ := os.Getwd()
	return strings.TrimRight(baseDir, "/tests")
}

// getEnv - Функция для получения значения переменной окружения.
// Принимает название переменной окружения string и значение по умолчанию string. Возвращает string
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// newPathConfig - Функция конструктор для получения экземпляра структуры PathConfig.
// Возвращает *PathConfig
func newPathConfig() *PathConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не удалось обнаружить файл .env", err)
	}

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

// ReadCsvFile - Функция для чтения csv файла. Принимает путь файла string. Возвращает содерижимое в формате [][]string
func ReadCsvFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Не удалось открыть файл", err.Error())
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Не удалось прочитать файл", err.Error())
	}
	return data
}

// GetAlpha2Code - Функция для получения Alpha2 кодов из файла. Принимает путь файла string.
// Возвращает содерижимое в формате []string
func GetAlpha2Code(path string) []string {
	alphaCode := make([]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		alphaCode = append(alphaCode, line[1])
	}
	return alphaCode
}

// GetCountryAlpha2Code - Функция для получения стран с Alpha2 кодом из файла. Принимает путь файла string.
// Возвращает содерижимое в формате map[string]string
func GetCountryAlpha2Code(path string) map[string]string {
	result := make(map[string]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		result[line[1]] = line[0]
	}
	return result
}

// GetAllowProviders - Функция для получения разрещенных провайдеров из файла. Принимает путь файла string.
// Возвращает содерижимое в формате []string
func GetAllowProviders(path string) []string {
	allowProviders := make([]string, 0)
	data := ReadCsvFile(path)
	for _, line := range data {
		allowProviders = append(allowProviders, line[0])
	}
	return allowProviders
}

// IsExist - Функция проверяет наличие элемента в срезе. Принимает []string и искомый элемент string. Возвращает bool.
func IsExist(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

// ToInt - Функция конвертирует string в int.
func ToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	return number
}

// ToFloat32 - Функция конвертирует string в float32.
func ToFloat32(str string) float32 {
	number, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Println(err)
	}
	return float32(number)
}

// ConvertToBool - Функция конвертирует byte в bool
func ConvertToBool(b byte) bool {
	if b == 48 {
		return false
	}
	return true
}

// ErrorToString - Функция конвертирует error в string
func ErrorToString(err ...error) string {
	var errorString string
	for _, item := range err {
		if item != nil {
			errorString += item.Error() + ", "
		}

	}
	return strings.TrimRight(errorString, ", ")
}
