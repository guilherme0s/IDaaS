package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/guilherme0s/crm/pkg/config"
)

type HTTPServer struct {
	cfg     config.Config
	httpSvr *http.Server
}

func NewHTTPServer(cfg config.Config) *HTTPServer {
	return &HTTPServer{
		cfg: cfg,
	}
}

func (hs *HTTPServer) Run(ctx context.Context) error {
	addr := net.JoinHostPort(hs.cfg.Host, hs.cfg.Port)

	hs.httpSvr = &http.Server{
		Addr: addr,
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to bind address: %w", err)
	}

	slog.Info("HTTP server listen", "address", listener.Addr().String())

	var wg sync.WaitGroup

	wg.Go(func() {

		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := hs.httpSvr.Shutdown(shutdownCtx); err != nil {
			slog.Error("failed to shutdown server", "error", err)
		}
	})

	if err := hs.httpSvr.Serve(listener); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("server shutdown gracefully")
			wg.Wait()
			return nil
		}
		return err
	}

	wg.Wait()
	return nil
}
