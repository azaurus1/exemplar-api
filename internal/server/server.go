package server

import (
	"exemplar-api/internal/config"
	"net/http"
)

type Server struct {
	Server *http.Server
	Config config.Config
	StopCh chan struct{}
}
