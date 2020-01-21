package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App Makes the web interface testable.
type App struct {
	Router *mux.Router
}

// Initialize initialises routing
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.helloHandler).Methods("GET")
}

func (a *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler helloHandler request received")

	w.WriteHeader(http.StatusOK)
	log.Printf("%s status [%d]\n", r.RequestURI, http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, `Hello There Gopher!`)
}

// Run runs the web server
func (a *App) Run(port string) {
	addr := fmt.Sprintf(":%s", port)
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Running action server on %s\n", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func main() {
	a := App{}
	a.Initialize()
	a.Run("8080")
}
