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
		DataTest{"1",
			[]sms_service.SMSData{
				{"FR", "68", "1049", "Topolo"},
				{"GB", "94", "1314", "Topolo"},
				{"CA", "8", "1420", "Rond"}},
			[][]sms_service.SMSData{
				[]sms_service.SMSData{{"Canada", "8", "1420", "Rond"}, {"France", "68", "1049", "Topolo"}, {"United Kingdom", "94", "1314", "Topolo"}},
				[]sms_service.SMSData{{"Canada", "8", "1420", "Rond"}, {"United Kingdom", "94", "1314", "Topolo"}, {"France", "68", "1049", "Topolo"}},
			}},
	}

	for _, test := range dataTest {
		assert.Equal(t, sms_service.SortedSmsData(test.input), test.expected)
	}
}
