package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kamijooou/L0.service/internal/cache"
	"github.com/kamijooou/L0.service/pkg/log"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type ServerHandler struct {
	srv *http.Server
}

func NewServer() *ServerHandler {
	srv := &http.Server{}
	srv.Addr = ":8001"

	router := mux.NewRouter()
	router.HandleFunc("/id/{uid:.*}", idHandler).Methods("GET")
	srv.Handler = router

	return &ServerHandler{srv: srv}
}

func (s *ServerHandler) Run(ctx context.Context) {
	logger := log.LoggerFromContext(ctx)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			logger.Error("server listen fail:", zap.Error(err))
			return
		}
	}()

	logger.Sugar().Infof("server start listening at localhost%s...", s.srv.Addr)
}

func (s *ServerHandler) Stop(ctx context.Context) {
	logger := log.LoggerFromContext(ctx)
	if err := s.srv.Shutdown(ctx); err != nil {
		logger.Error("server stop error:", zap.Error(err))
		return
	}
	logger.Info("server stopped")
}

func idHandler(rw http.ResponseWriter, req *http.Request) {
	logger, _ := log.NewLogger()
	ctx := log.ContextWithLogger(context.Background(), logger)

	params := strings.Split(req.URL.Path, "/")
	if len(params) != 3 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := cache.Get(ctx, params[2])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(*order)
	if err != nil {
		logger.Error("JSON marshalling:", zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("GET order:", zap.String("UID", params[2]))
	rw.Write(result)
	rw.WriteHeader(http.StatusOK)
}
