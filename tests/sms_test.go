package tests

import (
	sms_service "Diplom_Makarov/internal/sms-service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSmsData(t *testing.T) {

	type DataTest struct {
		name     string
		input    []sms_service.SMSData
		expected [][]sms_service.SMSData
	}

	var dataTest = []DataTest{
		DataTest{"Sorted_OK",
			[]sms_service.SMSData{
				{"FR", "68", "1049", "Topolo"},
				{"GB", "94", "1314", "Topolo"},
				{"CA", "8", "1420", "Rond"}},
			[][]sms_service.SMSData{
				{{"Canada", "8", "1420", "Rond"}, {"France", "68", "1049", "Topolo"}, {"United Kingdom", "94", "1314", "Topolo"}},
				{{"Canada", "8", "1420", "Rond"}, {"France", "68", "1049", "Topolo"}, {"United Kingdom", "94", "1314", "Topolo"}},
			}},
	}

	for _, test := range dataTest {
		assert.Equal(t, sms_service.SortedSmsData(test.input), test.expected)
	}
}

func TestValidateSmsData(t *testing.T) {

	type DataTest struct {
		name     string
		input    [][]string
		expected []sms_service.SMSData
	}

	var dataTest = []DataTest{
		{"OK_All",
			[][]string{{"RU;45;1208;Topolo"}, {"NZ;64;295;Kildy"}, {"GB;15;821;Topolo"}},
			[]sms_service.SMSData{
				{"RU", "45", "1208", "Topolo"},
				{"NZ", "64", "295", "Kildy"},
				{"GB", "15", "821", "Topolo"}},
		},
		{"OK_without_1_country",
			[][]string{{"ZZ;45;1208;Topolo"}, {"NZ;64;295;Kildy"}, {"GB;15;821;Topolo"}},
			[]sms_service.SMSData{
				{"NZ", "64", "295", "Kildy"},
				{"GB", "15", "821", "Topolo"}},
		},
		{"OK_without_1_provider",
			[][]string{{"RU;45;1208;Topolo"}, {"NZ;64;295;Kildy"}, {"GB;15;821;Topolox"}},
			[]sms_service.SMSData{
				{"RU", "45", "1208", "Topolo"},
				{"NZ", "64", "295", "Kildy"}},
		},
		{"OK_without_2_len",
			[][]string{{"RU;45;1208;Topolo;L"}, {"Kildy"}, {"GB;15;821;Topolo"}},
			[]sms_service.SMSData{
				{"GB", "15", "821", "Topolo"}},
		},
	}

	for _, test := range dataTest {
		assert.Equal(t, sms_service.ValidateSmsData(test.input), test.expected)
	}
}
