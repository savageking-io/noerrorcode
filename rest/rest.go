package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	restapi "github.com/savageking-io/noerrorcode/rest/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Config struct {
	Hostname       string   `yaml:"hostname"`
	Port           uint16   `yaml:"port"`
	AllowedOrigins []string `yaml:"allowed_origins"`
}

type REST struct {
	Hostname       string
	Port           uint16
	AllowedOrigins []string
	mux            *chi.Mux
}

func (d *REST) Init(config *Config) error {
	log.Traceln("REST::Init")

	log.Debugf("Config: %+v", config)
	d.Hostname = config.Hostname
	d.Port = config.Port
	d.AllowedOrigins = config.AllowedOrigins
	if len(d.AllowedOrigins) == 0 {
		d.AllowedOrigins = []string{"http://localhost:3000"}
	}
	return nil
}

func (d *REST) Start() error {
	log.Traceln("REST::Start")

	api := restapi.API{}

	d.mux = chi.NewMux()
	d.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   d.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	handler := restapi.HandlerFromMux(api, d.mux)

	log.Infof("Starting REST server on port %d", d.Port)
	s := &http.Server{Handler: handler, Addr: fmt.Sprintf(":%d", d.Port)}
	return s.ListenAndServe()
}
