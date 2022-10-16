package main

import (
	billing_service "Diplom_Makarov/internal/billing-service"
	email_service "Diplom_Makarov/internal/email-service"
	accendent_service "Diplom_Makarov/internal/incident-service"
	mms_service "Diplom_Makarov/internal/mms-service"
	sms_service "Diplom_Makarov/internal/sms-service"
	support_service "Diplom_Makarov/internal/support-service"
	"Diplom_Makarov/internal/voicecall-service"
)

func main() {
	sms_service.StartSmsService()
	mms_service.StartMmsService()
	voicecall_service.StartVoiceService()
	email_service.StartEmailService()
	billing_service.StartBillingService()
	support_service.StartSupportService()
	accendent_service.StartIncidentService()
}

//- Напишите функцию, которая будет читать всё содержимое из файла, далее  обходить содержимое построчно и разбирать строки на показатели
