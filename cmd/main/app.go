package main

import (
	"Diplom_Makarov/internal/parser-service"
	"Diplom_Makarov/internal/server"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не удалось обнаружить файл .env", err)
	}
}

func main() {
	parser_service.GetResultData()

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
