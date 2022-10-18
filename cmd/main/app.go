package main

import (
	parser_service "Diplom_Makarov/internal/parser-service"
)

func main() {
	parser_service.GetResultData()

	//router := mux.NewRouter()
	//handler := server.NewHandler()
	//handler.Register(router)
	//
	//serv := &http.Server{Addr: "localhost:8585", Handler: router}
	//go func() {
	//	if err := serv.ListenAndServe(); err != nil {
	//		log.Fatal("не удалось запустить сервер: ", err)
	//	}
	//}()
	//
	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//<-stop
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//serv.Shutdown(ctx)
}
