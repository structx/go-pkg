package serverfx

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/trevatk/go-pkg/domain"
)

// New http server provider
func New(cfg domain.Config, handler http.Handler) *http.Server {
	scfg := cfg.GetServer()
	return &http.Server{
		Addr:         net.JoinHostPort(scfg.BindAddr, fmt.Sprintf("%d", scfg.Ports.HTTP)),
		Handler:      handler,
		ReadTimeout:  time.Duration(scfg.DefaultTimeout) * time.Second,
		WriteTimeout: time.Duration(scfg.DefaultTimeout) * time.Second,
		IdleTimeout:  time.Duration(scfg.DefaultTimeout) * time.Second,
	}
}