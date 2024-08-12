package webserver

import (
	"log"
	"net/http"
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
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	muxHandler := cors(authMiddleware(s.Router))

	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	log.Println("Starting web server on port", s.WebServerPort)
	err := http.ListenAndServe(s.WebServerPort, muxHandler)
	if err != nil {
		panic(err)
	}
}
