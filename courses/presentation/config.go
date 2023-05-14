package presentation

import (
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MelvinKim/courses/infrastructure/database"
	"github.com/MelvinKim/courses/presentation/interactor"
	"github.com/MelvinKim/courses/presentation/rest"
	"github.com/MelvinKim/courses/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	serverTimeoutSeconds = 120
)

var allowedHeaders = []string{
	"Authorization", "Accept", "Accept-Charset", "Accept-Language",
	"Accept-Encoding", "Origin", "Host", "User-Agent", "Content-Length",
	"Content-Type", " X-Authorization", " Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers",
}

// Router sets up the gorilla Mux router
func Router(ctx context.Context) (*mux.Router, error) {
	create := database.NewPostgresDB()
	get := database.NewPostgresDB()
	users := usecase.NewUsecase(create, get)

	i, err := interactor.NewUsersInteractor(
		users,
	)
	if err != nil {
		return nil, fmt.Errorf("can't instantiate a new service: %w", err)
	}

	h := rest.NewPresentationHandlers(i)

	r := mux.NewRouter()

	userRoutes := r.PathPrefix("/api/v1").Subrouter()
	userRoutes.Path("/users").Methods(http.MethodPost).HandlerFunc(h.CreateStudent())
	userRoutes.Path("/user").Methods(http.MethodGet).HandlerFunc(h.GetStudent())
	userRoutes.Path("/courses").Methods(http.MethodPost).HandlerFunc(h.CreateCourse())
	userRoutes.Path("/course").Methods(http.MethodGet).HandlerFunc(h.GetCourse())
	userRoutes.Path("/assign_course").Methods(http.MethodPost).HandlerFunc(h.AssignCourseToStudent())

	return r, nil
}

// PrepareServer starts up a server
func PrepareServer(
	ctx context.Context,
	port int,
) *http.Server {
	// start up  the router
	r, err := Router(ctx)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Server startup error")
	}

	// start the server
	addr := fmt.Sprintf(":%d", port)
	h := handlers.CompressHandlerLevel(r, gzip.BestCompression)

	h = handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST"}),
	)(h)
	h = handlers.CombinedLoggingHandler(os.Stdout, h)
	h = handlers.ContentTypeHandler(
		h,
		"application/json",
		"application/x-www-form-urlencoded",
	)
	srv := &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: serverTimeoutSeconds * time.Second,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
	}
	log.Infof("Server running at port %v", addr)
	return srv

}
