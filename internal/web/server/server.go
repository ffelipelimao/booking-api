package web

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type WebServer struct {
	Router        *fiber.App
	Handlers      map[string]func(h *fiber.Ctx) error
	WebServerPort string
	Verb          string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        fiber.New(),
		Handlers:      make(map[string]func(h *fiber.Ctx) error),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, handler func(h *fiber.Ctx) error, verb string) {
	s.Handlers[path] = handler
	s.Verb = verb
}

func (s *WebServer) Start() {
	s.Router.Use(logger.New())

	for path, handler := range s.Handlers {
		switch s.Verb {
		case "POST":
			s.Router.Post(path, handler)
		case "GET":
			s.Router.Get(path, handler)
		case "PUT":
			s.Router.Put(path, handler)
		default:
			log.Panic("invalid http verb")
		}
	}

	s.Router.Listen(":3000")
}
