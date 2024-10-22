package webserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type WebServer struct {
	Router        *http.ServeMux
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        http.NewServeMux(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	if s.Handlers[path] != nil {
		panic("identical http routes cannot be created")
	}

	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	muxHandler := cors(authMiddleware(s.Router))

	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	server := &http.Server{
		Addr:    s.WebServerPort,
		Handler: muxHandler,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Starting web server on port", s.WebServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", s.WebServerPort, err)
		}
	}()

	<-stop

	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
