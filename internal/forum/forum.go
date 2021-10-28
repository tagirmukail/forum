package forum

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/tagirmukail/forum/docs"
	"github.com/tagirmukail/forum/internal/config"
	"github.com/tagirmukail/forum/internal/forum/logic"
	"github.com/tagirmukail/forum/internal/repository"
)

const (
	contentTypeJSON = "application/json"

	requestBodyDeserializationFailed = "request body deserialization failed"
	respSerializationFailed          = "response serialization failed"
	invalidLimitParam                = "invalid limit parameter"
	invalidOffsetParam               = "invalid offset parameter"
)

type Dependencies struct {
	Repo repository.Repository
	Conf *config.Config
}

type Service struct {
	d *Dependencies
	l *logic.Logic
}

func NewService(d *Dependencies) *Service {
	return &Service{
		d: d,
		l: logic.NewLogic(&logic.Dependencies{
			Repo: d.Repo,
			Conf: d.Conf,
		}),
	}
}

func (s *Service) Serve() error {
	serv := &http.Server{
		Handler:      corsHandler()(s.router()),
		Addr:         s.d.Conf.API.Addr,
		WriteTimeout: time.Duration(s.d.Conf.API.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(s.d.Conf.API.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.d.Conf.API.IdleTimeout) * time.Second,
	}

	return serv.ListenAndServe()
}

func (s *Service) router() *mux.Router {
	r := mux.NewRouter()

	r.Use(
		s.requestContentLimitMiddleware(),
		s.requestIDMiddleware(),
		s.loggingMiddleware(),
	)

	r.HandleFunc("/api/v1/users", s.createUser).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/topics", s.listTopics).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/topics/{topic_id}", s.getTopic).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/topics", s.createTopic).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/topics/{topic_id}/comments", s.listComments).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/topics/{topic_id}/comments", s.createComment).Methods(http.MethodPost)

	r.HandleFunc("/health", s.health).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DocExpansion("none"),
	))

	return r
}

func (s *Service) setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type HealthResponse struct {
	Data string `json:"data" example:"Service is up and running"`
}

// @Summary Health
// @Description service health check
// @Tags Forum
// @Produce json
// @Success 200 {object} HealthResponse "Success operation"
// @Router /health [get]
//
func (s *Service) health(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	w.Header().Set("Content-Type", contentTypeJSON)

	b, err := jsoniter.Marshal(HealthResponse{
		Data: "Service is up and running",
	})
	if err != nil {
		wr.ERR = err
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = w.Write(b)
	if err != nil {
		wr.ERR = err
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
