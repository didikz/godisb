package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/didikz/godisb/api/handler"
	"github.com/didikz/godisb/config"
	"github.com/didikz/godisb/pkg/db"
	m "github.com/didikz/godisb/pkg/httpapi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

type HttpServer struct {
	listenAddress string
	db            *sqlx.DB
	cfg           config.Configuration
}

func NewHttpServer(cfg config.Configuration) *HttpServer {
	dbx := db.NewDB(db.ConfigDB{
		Driver:   cfg.DB.Driver,
		Name:     cfg.DB.Name,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
	})
	return &HttpServer{
		listenAddress: fmt.Sprintf(":%s", cfg.App.Port),
		db:            dbx,
		cfg:           cfg,
	}
}

func (s *HttpServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(m.HeaderValidator)
	handler.RegisterHandler(r, s.db, s.cfg)

	log.Println("running server at", s.listenAddress)
	return http.ListenAndServe(s.listenAddress, r)
}
