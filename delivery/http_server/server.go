package httpserver

import (
	"fmt"
	"github.com/amiranbari/challenge/config"
	"github.com/amiranbari/challenge/service"
	echo "github.com/labstack/echo/v4"
)

type Server struct {
	config config.Config
	Router *echo.Echo
}

func New(
	cfg config.Config,
	svc *service.Service,
) *Server {
	return &Server{
		Router: echo.New(),
		config: cfg,
	}
}

func (s Server) Serve() {
	s.RegisterRoutes()

	// Start server
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}

func (s Server) RegisterRoutes() {
	// Routes
}
