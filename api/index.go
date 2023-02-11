package main

import (
	"arrow_food_api/handler"
	handlerlogin "arrow_food_api/handler/handlerLogin"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()


	 /* mail :=  models.EmailTemplate{Sender: "axcefornax@gmail.com",Subject: "prueba email",Body: "este es un email de prueba"}

	 sendMail := models.Email{}

	 sendMail.SendMail(&mail) */



	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"hola mundo...")
	}).Methods("GET")

	mux.HandleFunc("/api/users/{email}", handler.GetUserByEmail).Methods("GET").Name("users")
	mux.HandleFunc("/api/login", handlerlogin.Login).Methods("POST")
	mux.HandleFunc("/api/register/users", handler.RegisterUser).Methods("POST")
	mux.HandleFunc("/api/users/delete/{email}", handler.DeleteUser).Methods("DELETE")

	server := &http.Server{
		Handler: mux,
		Addr: ":3000",
		WriteTimeout: 15* time.Second,
		ReadTimeout: 15*time.Second,
	}

	go func() {
		fmt.Println("server is listening on " + server.Addr)
		err := server.ListenAndServe();
		if err != nil {
			log.Println(err)
		}
	}()

	var wait time.Duration

	c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <- c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    server.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)
}