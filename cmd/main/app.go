package main

import (
	billing_service "Diplom_Makarov/internal/billing-service"
	email_service "Diplom_Makarov/internal/email-service"
	accendent_service "Diplom_Makarov/internal/incident-service"
	mms_service "Diplom_Makarov/internal/mms-service"
	"Diplom_Makarov/internal/server"
	sms_service "Diplom_Makarov/internal/sms-service"
	support_service "Diplom_Makarov/internal/support-service"
	"Diplom_Makarov/internal/voicecall-service"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sms_service.StartSmsService()
	mms_service.StartMmsService()
	voicecall_service.StartVoiceService()
	email_service.StartEmailService()
	billing_service.StartBillingService()
	support_service.StartSupportService()
	accendent_service.StartIncidentService()

	router := mux.NewRouter()
	handler := server.NewHandler()
	handler.Register(router)

	serv := &http.Server{Addr: "localhost:8585", Handler: router}
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Fatal("не удалось запустить сервер: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serv.Shutdown(ctx)
}
