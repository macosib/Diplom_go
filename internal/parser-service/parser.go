package parser_service

import (
	billing_service "Diplom_Makarov/internal/billing-service"
	email_service "Diplom_Makarov/internal/email-service"
	incident_service "Diplom_Makarov/internal/incident-service"
	mms_service "Diplom_Makarov/internal/mms-service"
	sms_service "Diplom_Makarov/internal/sms-service"
	support_service "Diplom_Makarov/internal/support-service"
	voicecall_service "Diplom_Makarov/internal/voicecall-service"
)

type ResultT struct {
	Status bool        `json:"status"` // true, если все этапы сбора  данных прошли успешно, false во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string      `json:"error"`  // пустая строка если все этапы  сбора данных прошли успешно, в случае ошибки заполнено
	//текстом ошибки (детали ниже)
}

type ResultSetT struct {
	SMS       [][]sms_service.SMSData                `json:"sms"`
	MMS       [][]mms_service.MMSData                `json:"mms"`
	VoiceCall []voicecall_service.VoiceData          `json:"voice_call"`
	Email     map[string][][]email_service.EmailData `json:"email"`
	Billing   billing_service.BillingData            `json:"billing"`
	Support   []int                                  `json:"support"`
	Incidents []incident_service.IncidentData        `json:"incident"`
}

func GetResultData() ResultT {

	sms := sms_service.StartSmsService()
	mms, errMms := mms_service.StartMmsService()
	voice := voicecall_service.StartVoiceService()
	email := email_service.StartEmailService()
	billing := billing_service.StartBillingService()
	support, errSupport := support_service.StartSupportService()
	incident, errIncident := incident_service.StartIncidentService()

	if errMms != nil || errSupport != nil || errIncident != nil {
		return ResultT{false, nil, errMms.Error() + errSupport.Error() + errIncident.Error()}
	}

	return ResultT{true, &ResultSetT{sms, mms, voice, email, *billing, support, incident}, ""}
	//
	//
	//
	//fmt.Println(sms)
	//fmt.Println(mms)
	//fmt.Println(voice)
	//fmt.Println(email)
	//fmt.Println(billing)
	//fmt.Println(support)
	//fmt.Println(incident)
	//fmt.Println(errMms)
	//fmt.Println(errSupport)
	//fmt.Println(errIncident)

}
