package parser_service

import (
	billing_service "Diplom_Makarov/internal/billing-service"
	email_service "Diplom_Makarov/internal/email-service"
	incident_service "Diplom_Makarov/internal/incident-service"
	mms_service "Diplom_Makarov/internal/mms-service"
	sms_service "Diplom_Makarov/internal/sms-service"
	support_service "Diplom_Makarov/internal/support-service"
	"Diplom_Makarov/internal/utils"
	voicecall_service "Diplom_Makarov/internal/voicecall-service"
)

// ResultT - Результирующая структура с данными для передачи по HTTP
type ResultT struct {
	Status bool        `json:"status"`
	Data   *ResultSetT `json:"data"`
	Error  string      `json:"error"`
}

// ResultSetT - Структура для хранениния данных со всех сервисов
type ResultSetT struct {
	SMS       [][]sms_service.SMSData                `json:"sms"`
	MMS       [][]mms_service.MMSData                `json:"mms"`
	VoiceCall []voicecall_service.VoiceData          `json:"voice_call"`
	Email     map[string][][]email_service.EmailData `json:"email"`
	Billing   billing_service.BillingData            `json:"billing"`
	Support   []int                                  `json:"support"`
	Incidents []incident_service.IncidentData        `json:"incident"`
}

// GetResultData - Функция собирает данные со всех сервисов и возвращает результ в виде структуры ResultT
func GetResultData() ResultT {
	sms := sms_service.StartSmsService()
	mms, errMms := mms_service.StartMmsService()
	voice := voicecall_service.StartVoiceService()
	email := email_service.StartEmailService()
	billing := billing_service.StartBillingService()
	support, errSupport := support_service.StartSupportService()
	incident, errIncident := incident_service.StartIncidentService()

	if errMms != nil || errSupport != nil || errIncident != nil {
		return ResultT{false, nil, utils.ErrorToString(errMms, errSupport, errIncident)}
	}

	return ResultT{true, &ResultSetT{sms, mms, voice, email, *billing, support, incident}, ""}
}
