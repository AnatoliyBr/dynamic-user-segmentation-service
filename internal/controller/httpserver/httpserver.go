package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	uc     usecase.UseCase
}

func NewServer(config *Config, uc usecase.UseCase) *server {
	s := &server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		uc:     uc,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello()).Methods(http.MethodGet)

	s.router.HandleFunc("/seg", s.handleSegmentsCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/seg", s.handleSegmentsDelete()).Methods(http.MethodDelete)
	//s.router.HandleFunc("/seg", s.handleSegmentsUpdateUser()).Methods(http.MethodPut)
	//s.router.HandleFunc("/seg", s.handleSegmentsGetByUser()).Methods(http.MethodGet)
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *server) StartServer() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting http server")

	return http.ListenAndServe(s.config.BindAddr, s)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]string{"test": "hello"})
	}
}

func (s *server) handleSegmentsCreate() http.HandlerFunc {
	type request struct {
		Slug string `json:"slug"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		seg := &entity.Segment{
			Slug: req.Slug,
		}

		if err := s.uc.SegmentCreate(seg); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, seg)
	}
}

func (s *server) handleSegmentsDelete() http.HandlerFunc {
	type request struct {
		Slug string `json:"slug"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		seg, err := s.uc.SegmentFindBySlug(req.Slug)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		if err := s.uc.SegmentDelete(seg); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusOK, map[string]string{"delete segment": seg.Slug})
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
